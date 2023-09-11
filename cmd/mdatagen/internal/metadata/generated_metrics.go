// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

// AttributeEnumAttr specifies the a value enum_attr attribute.
type AttributeEnumAttr int

const (
	_ AttributeEnumAttr = iota
	AttributeEnumAttrRed
	AttributeEnumAttrGreen
	AttributeEnumAttrBlue
)

// String returns the string representation of the AttributeEnumAttr.
func (av AttributeEnumAttr) String() string {
	switch av {
	case AttributeEnumAttrRed:
		return "red"
	case AttributeEnumAttrGreen:
		return "green"
	case AttributeEnumAttrBlue:
		return "blue"
	}
	return ""
}

// MapAttributeEnumAttr is a helper map of string to AttributeEnumAttr attribute value.
var MapAttributeEnumAttr = map[string]AttributeEnumAttr{
	"red":   AttributeEnumAttrRed,
	"green": AttributeEnumAttrGreen,
	"blue":  AttributeEnumAttrBlue,
}

type metricDefaultMetric struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills default.metric metric with initial data.
func (m *metricDefaultMetric) init() {
	m.data.SetName("default.metric")
	m.data.SetDescription("Monotonic cumulative sum int metric enabled by default.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDefaultMetric) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, stringAttrAttributeValue string, overriddenIntAttrAttributeValue int64, enumAttrAttributeValue string, sliceAttrAttributeValue []any, mapAttrAttributeValue map[string]any) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("string_attr", stringAttrAttributeValue)
	dp.Attributes().PutInt("state", overriddenIntAttrAttributeValue)
	dp.Attributes().PutStr("enum_attr", enumAttrAttributeValue)
	dp.Attributes().PutEmptySlice("slice_attr").FromRaw(sliceAttrAttributeValue)
	dp.Attributes().PutEmptyMap("map_attr").FromRaw(mapAttrAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDefaultMetric) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDefaultMetric) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDefaultMetric(cfg MetricConfig) metricDefaultMetric {
	m := metricDefaultMetric{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDefaultMetricToBeRemoved struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills default.metric.to_be_removed metric with initial data.
func (m *metricDefaultMetricToBeRemoved) init() {
	m.data.SetName("default.metric.to_be_removed")
	m.data.SetDescription("[DEPRECATED] Non-monotonic delta sum double metric enabled by default.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityDelta)
}

func (m *metricDefaultMetricToBeRemoved) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDefaultMetricToBeRemoved) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDefaultMetricToBeRemoved) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDefaultMetricToBeRemoved(cfg MetricConfig) metricDefaultMetricToBeRemoved {
	m := metricDefaultMetricToBeRemoved{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricOptionalHistogramMetric struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills optional.histogram.metric metric with initial data.
func (m *metricOptionalHistogramMetric) init() {
	m.data.SetName("optional.histogram.metric")
	m.data.SetDescription("[DEPRECATED] Histogram metric disabled by default.")
	m.data.SetUnit("1")
	m.data.SetEmptyHistogram()
	m.data.Histogram().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Histogram().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricOptionalHistogramMetric) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, count uint64, sum float64, bucketCounts []uint64, explicitBounds []float64, stringAttrAttributeValue string, booleanAttrAttributeValue bool) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Histogram().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetCount(count)
	dp.SetSum(sum)
	if len(bucketCounts) > 0 {
		dp.BucketCounts().FromRaw(bucketCounts)
	}
	if len(explicitBounds) > 0 {
		dp.ExplicitBounds().FromRaw(explicitBounds)
	}
	dp.Attributes().PutStr("string_attr", stringAttrAttributeValue)
	dp.Attributes().PutBool("boolean_attr", booleanAttrAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricOptionalHistogramMetric) updateCapacity() {
	if m.data.Histogram().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Histogram().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricOptionalHistogramMetric) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Histogram().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricOptionalHistogramMetric(cfg MetricConfig) metricOptionalHistogramMetric {
	m := metricOptionalHistogramMetric{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricOptionalMetric struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills optional.metric metric with initial data.
func (m *metricOptionalMetric) init() {
	m.data.SetName("optional.metric")
	m.data.SetDescription("[DEPRECATED] Gauge double metric disabled by default.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricOptionalMetric) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, stringAttrAttributeValue string, booleanAttrAttributeValue bool) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("string_attr", stringAttrAttributeValue)
	dp.Attributes().PutBool("boolean_attr", booleanAttrAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricOptionalMetric) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricOptionalMetric) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricOptionalMetric(cfg MetricConfig) metricOptionalMetric {
	m := metricOptionalMetric{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                         MetricsBuilderConfig // config of the metrics builder.
	startTime                      pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                int                  // maximum observed number of metrics per resource.
	metricsBuffer                  pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                      component.BuildInfo  // contains version information.
	metricDefaultMetric            metricDefaultMetric
	metricDefaultMetricToBeRemoved metricDefaultMetricToBeRemoved
	metricOptionalHistogramMetric  metricOptionalHistogramMetric
	metricOptionalMetric           metricOptionalMetric
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	if !mbc.Metrics.DefaultMetric.enabledSetByUser {
		settings.Logger.Warn("[WARNING] Please set `enabled` field explicitly for `default.metric`: This metric will be disabled by default soon.")
	}
	if mbc.Metrics.DefaultMetricToBeRemoved.Enabled {
		settings.Logger.Warn("[WARNING] `default.metric.to_be_removed` should not be enabled: This metric is deprecated and will be removed soon.")
	}
	if mbc.Metrics.OptionalHistogramMetric.enabledSetByUser {
		settings.Logger.Warn("[WARNING] `optional.histogram.metric` should not be configured: This metric is deprecated and will be removed soon.")
	}
	if mbc.Metrics.OptionalMetric.enabledSetByUser {
		settings.Logger.Warn("[WARNING] `optional.metric` should not be configured: This metric is deprecated and will be removed soon.")
	}
	mb := &MetricsBuilder{
		config:                         mbc,
		startTime:                      pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                  pmetric.NewMetrics(),
		buildInfo:                      settings.BuildInfo,
		metricDefaultMetric:            newMetricDefaultMetric(mbc.Metrics.DefaultMetric),
		metricDefaultMetricToBeRemoved: newMetricDefaultMetricToBeRemoved(mbc.Metrics.DefaultMetricToBeRemoved),
		metricOptionalHistogramMetric:  newMetricOptionalHistogramMetric(mbc.Metrics.OptionalHistogramMetric),
		metricOptionalMetric:           newMetricOptionalMetric(mbc.Metrics.OptionalMetric),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted metrics.
func (mb *MetricsBuilder) NewResourceBuilder() *ResourceBuilder {
	return NewResourceBuilder(mb.config.ResourceAttributes)
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricDefaultMetric.emit(ils.Metrics())
	mb.metricDefaultMetricToBeRemoved.emit(ils.Metrics())
	mb.metricOptionalHistogramMetric.emit(ils.Metrics())
	mb.metricOptionalMetric.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordDefaultMetricDataPoint adds a data point to default.metric metric.
func (mb *MetricsBuilder) RecordDefaultMetricDataPoint(ts pcommon.Timestamp, val int64, stringAttrAttributeValue string, overriddenIntAttrAttributeValue int64, enumAttrAttributeValue AttributeEnumAttr, sliceAttrAttributeValue []any, mapAttrAttributeValue map[string]any) {
	mb.metricDefaultMetric.recordDataPoint(mb.startTime, ts, val, stringAttrAttributeValue, overriddenIntAttrAttributeValue, enumAttrAttributeValue.String(), sliceAttrAttributeValue, mapAttrAttributeValue)
}

// RecordDefaultMetricToBeRemovedDataPoint adds a data point to default.metric.to_be_removed metric.
func (mb *MetricsBuilder) RecordDefaultMetricToBeRemovedDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricDefaultMetricToBeRemoved.recordDataPoint(mb.startTime, ts, val)
}

// RecordOptionalHistogramMetricDataPoint adds a data point to optional.histogram.metric metric.
func (mb *MetricsBuilder) RecordOptionalHistogramMetricDataPoint(ts pcommon.Timestamp, count uint64, sum float64, bucketCounts []uint64, explicitBounds []float64, stringAttrAttributeValue string, booleanAttrAttributeValue bool) {
	mb.metricOptionalHistogramMetric.recordDataPoint(mb.startTime, ts, count, sum, bucketCounts, explicitBounds, stringAttrAttributeValue, booleanAttrAttributeValue)
}

// RecordOptionalMetricDataPoint adds a data point to optional.metric metric.
func (mb *MetricsBuilder) RecordOptionalMetricDataPoint(ts pcommon.Timestamp, val float64, stringAttrAttributeValue string, booleanAttrAttributeValue bool) {
	mb.metricOptionalMetric.recordDataPoint(mb.startTime, ts, val, stringAttrAttributeValue, booleanAttrAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
