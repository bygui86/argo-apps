package monitoring

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// WARN this object is not thread safe!
// customMetrics - Container object for custom metrics
type customMetrics struct {
	counters   map[string]prometheus.Counter
	gauges     map[string]prometheus.Gauge
	summaries  map[string]prometheus.Summary
	histograms map[string]prometheus.Histogram
}

type ICustomMetrics interface {
	AddCounter(namespace, subsystem, name, help, internalKey string)
	AddGauge(namespace, subsystem, name, help, internalKey string)
	AddSummary(
		namespace, subsystem, name, help, internalKey string,
		maxAge time.Duration, constLabels prometheus.Labels,
		objectives map[float64]float64, ageBuckets, bufCap uint32)
	AddHistogram(
		namespace, subsystem, name, help, internalKey string,
		constLabels prometheus.Labels, buckets []float64)
	IncreaseCounter(counter string)
	SetGauge(gauge string, value float64)
	ObserveSummary(summary string, observation float64)
	ObserveHistogram(histogram string, observation float64)
}

// newCustomMetrics - Create CustomMetrics object
func newCustomMetrics() ICustomMetrics {

	counters := make(map[string]prometheus.Counter)
	gauges := make(map[string]prometheus.Gauge)
	summaries := make(map[string]prometheus.Summary)
	histograms := make(map[string]prometheus.Histogram)

	return &customMetrics{
		counters:   counters,
		gauges:     gauges,
		summaries:  summaries,
		histograms: histograms,
	}
}

// AddCounter - Add a counter to custom metrics
func (cm *customMetrics) AddCounter(namespace, subsystem, name, help, internalKey string) {

	cm.counters[internalKey] = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
	})
}

// AddGauge - Add a gauge to custom metrics
func (cm *customMetrics) AddGauge(namespace, subsystem, name, help, internalKey string) {

	cm.gauges[internalKey] = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
	})
}

// AddSummary - Add a summary to custom metrics
func (cm *customMetrics) AddSummary(
	namespace, subsystem, name, help, internalKey string,
	maxAge time.Duration, constLabels prometheus.Labels,
	objectives map[float64]float64, ageBuckets, bufCap uint32) {

	cm.summaries[internalKey] = promauto.NewSummary(prometheus.SummaryOpts{
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		MaxAge:      maxAge,
		ConstLabels: constLabels,
		Objectives:  objectives,
		AgeBuckets:  ageBuckets,
		BufCap:      bufCap,
	})
}

// AddHistogram - Add a histogram to custom metrics
func (cm *customMetrics) AddHistogram(
	namespace, subsystem, name, help, internalKey string,
	constLabels prometheus.Labels, buckets []float64) {

	cm.histograms[internalKey] = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
		Buckets:     buckets,
	})
}

// IncreaseCounter - Increase a counter
func (cm *customMetrics) IncreaseCounter(counter string) {

	cm.counters[counter].Inc()
}

// SetGauge - Set a gauge to a new value
func (cm *customMetrics) SetGauge(gauge string, value float64) {

	cm.gauges[gauge].Set(value)
}

// ObserveSummary - Add a single observation to a summary
func (cm *customMetrics) ObserveSummary(summary string, observation float64) {

	cm.summaries[summary].Observe(observation)
}

// ObserveHistogram - Add a single observation to an histogram
func (cm *customMetrics) ObserveHistogram(histogram string, observation float64) {

	cm.histograms[histogram].Observe(observation)
}
