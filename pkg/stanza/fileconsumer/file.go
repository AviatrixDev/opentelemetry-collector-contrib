// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileconsumer // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer"

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/checkpoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fingerprint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

type Manager struct {
	*zap.SugaredLogger
	wg     sync.WaitGroup
	cancel context.CancelFunc

	readerFactory reader.Factory
	fileMatcher   *matcher.Matcher

	pollInterval  time.Duration
	persister     operator.Persister
	maxBatches    int
	maxBatchFiles int

	previousPollFiles []*reader.Reader
	knownFiles        []*reader.Metadata
}

func (m *Manager) Start(persister operator.Persister) error {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel

	if persister != nil {
		m.persister = persister
		offsets, err := checkpoint.Load(ctx, m.persister)
		if err != nil {
			return fmt.Errorf("read known files from database: %w", err)
		}
		if len(offsets) > 0 {
			m.Infow("Resuming from previously known offset(s). 'start_at' setting is not applicable.")
			m.readerFactory.FromBeginning = true
			m.knownFiles = append(m.knownFiles, offsets...)
		}
	}

	if _, err := m.fileMatcher.MatchFiles(); err != nil {
		m.Warnf("finding files: %v", err)
	}

	// Start polling goroutine
	m.startPoller(ctx)

	return nil
}

func (m *Manager) closePreviousFiles() {
	if forgetNum := len(m.previousPollFiles) + len(m.knownFiles) - cap(m.knownFiles); forgetNum > 0 {
		m.knownFiles = m.knownFiles[forgetNum:]
	}
	for _, r := range m.previousPollFiles {
		m.knownFiles = append(m.knownFiles, r.Close())
	}
}

// Stop will stop the file monitoring process
func (m *Manager) Stop() error {
	m.cancel()
	m.wg.Wait()
	m.closePreviousFiles()
	if m.persister != nil {
		if err := checkpoint.Save(context.Background(), m.persister, m.knownFiles); err != nil {
			m.Errorw("save offsets", zap.Error(err))
		}
	}
	m.cancel = nil
	return nil
}

// startPoller kicks off a goroutine that will poll the filesystem periodically,
// checking if there are new files or new logs in the watched files
func (m *Manager) startPoller(ctx context.Context) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		globTicker := time.NewTicker(m.pollInterval)
		defer globTicker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-globTicker.C:
			}

			m.poll(ctx)
		}
	}()
}

// poll checks all the watched paths for new entries
func (m *Manager) poll(ctx context.Context) {
	// Used to keep track of the number of batches processed in this poll cycle
	batchesProcessed := 0

	// Get the list of paths on disk
	matches, err := m.fileMatcher.MatchFiles()
	if err != nil {
		m.Debugf("finding files: %v", err)
	}
	m.Debugf("matched files", zap.Strings("paths", matches))

	for len(matches) > m.maxBatchFiles {
		m.consume(ctx, matches[:m.maxBatchFiles])

		// If a maxBatches is set, check if we have hit the limit
		if m.maxBatches != 0 {
			batchesProcessed++
			if batchesProcessed >= m.maxBatches {
				return
			}
		}

		matches = matches[m.maxBatchFiles:]
	}
	m.consume(ctx, matches)

	// Any new files that appear should be consumed entirely
	m.readerFactory.FromBeginning = true
	if m.persister != nil {
		if err := checkpoint.Save(context.Background(), m.persister, m.knownFiles); err != nil {
			m.Errorw("save offsets", zap.Error(err))
		}
	}
}

func (m *Manager) consume(ctx context.Context, paths []string) {
	m.Debug("Consuming files", zap.Strings("paths", paths))
	readers := m.makeReaders(paths)

	// take care of files which disappeared from the pattern since the last poll cycle
	// this can mean either files which were removed, or rotated into a name not matching the pattern
	// we do this before reading existing files to ensure we emit older log lines before newer ones
	m.readLostFiles(ctx, readers)
	m.closePreviousFiles()

	// read new readers to end
	var wg sync.WaitGroup
	for _, r := range readers {
		wg.Add(1)
		go func(r *reader.Reader) {
			defer wg.Done()
			r.ReadToEnd(ctx)
		}(r)
	}
	wg.Wait()

	m.previousPollFiles = readers
}

func (m *Manager) makeFingerprint(path string) (*fingerprint.Fingerprint, *os.File) {
	file, err := os.Open(path) // #nosec - operator must read in files defined by user
	if err != nil {
		m.Errorw("Failed to open file", zap.Error(err))
		return nil, nil
	}

	fp, err := m.readerFactory.NewFingerprint(file)
	if err != nil {
		if err = file.Close(); err != nil {
			m.Debugw("problem closing file", zap.Error(err))
		}
		return nil, nil
	}

	if len(fp.FirstBytes) == 0 {
		// Empty file, don't read it until we can compare its fingerprint
		if err = file.Close(); err != nil {
			m.Debugw("problem closing file", zap.Error(err))
		}
		return nil, nil
	}
	return fp, file
}

// makeReader take a file path, then creates reader,
// discarding any that have a duplicate fingerprint to other files that have already
// been read this polling interval
func (m *Manager) makeReaders(paths []string) []*reader.Reader {
	readers := make([]*reader.Reader, 0, len(paths))
	for _, path := range paths {
		fp, file := m.makeFingerprint(path)
		if fp == nil {
			continue
		}

		// Exclude duplicate paths with the same content. This can happen when files are
		// being rotated with copy/truncate strategy. (After copy, prior to truncate.)
		for _, r := range readers {
			if fp.Equal(r.Fingerprint) {
				if err := file.Close(); err != nil {
					m.Debugw("problem closing file", zap.Error(err))
				}
				continue
			}
		}

		r, err := m.newReader(file, fp)
		if err != nil {
			m.Errorw("Failed to create reader", zap.Error(err))
			continue
		}

		readers = append(readers, r)
	}
	return readers
}

func (m *Manager) newReader(file *os.File, fp *fingerprint.Fingerprint) (*reader.Reader, error) {
	// Check previous poll cycle for match
	for i := 0; i < len(m.previousPollFiles); i++ {
		oldReader := m.previousPollFiles[i]
		if fp.StartsWith(oldReader.Fingerprint) {
			// Keep the new reader and discard the old. This ensures that if the file was
			// copied to another location and truncated, our handle is updated.
			m.previousPollFiles = append(m.previousPollFiles[:i], m.previousPollFiles[i+1:]...)
			return m.readerFactory.NewReaderFromMetadata(file, oldReader.Close())
		}
	}

	// Iterate backwards to match newest first
	for i := len(m.knownFiles) - 1; i >= 0; i-- {
		oldMetadata := m.knownFiles[i]
		if fp.StartsWith(oldMetadata.Fingerprint) {
			// Remove the old metadata from the list. We will keep updating it and save it again later.
			m.knownFiles = append(m.knownFiles[:i], m.knownFiles[i+1:]...)
			return m.readerFactory.NewReaderFromMetadata(file, oldMetadata)
		}
	}

	// If we don't match any previously known files, create a new reader from scratch
	m.Infow("Started watching file", "path", file.Name())
	return m.readerFactory.NewReader(file, fp)
}
