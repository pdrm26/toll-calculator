package main

import (
	"time"

	"github.com/pdrm26/toll-calculator/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type MetricMiddleware struct {
	reqCounter prometheus.Counter
	errCounter prometheus.Counter
	reqLatency prometheus.Histogram
	next       Aggregator
}

func NewMetricMiddleware(next Aggregator) *MetricMiddleware {
	reqCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "aggregator_request_counter",
		Help: "The total number of HTTP requests processed by the aggregator service",
	})
	errCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "aggregator_error_counter",
		Help: "The total number of HTTP errors occur in the aggregator service",
	})
	reqLatency := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "aggregator_request_latency",
		Help:    "The latency of HTTP requests handled by the aggregator service, in seconds",
		Buckets: []float64{0.1, 0.5, 1}, // buckets for 100ms, 500ms, 1s
	})
	return &MetricMiddleware{
		reqCounter: reqCounter,
		errCounter: errCounter,
		reqLatency: reqLatency,
		next:       next,
	}
}

func (m *MetricMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		m.reqLatency.Observe(time.Since(start).Seconds())
		m.reqCounter.Inc()
		if err != nil {
			m.errCounter.Inc()
		}
	}(time.Now())
	err = m.next.AggregateDistance(distance)
	return
}

// TODO: write the metric for this endpoint also
func (m *MetricMiddleware) CalculateInvoice(obuID int) (invoice *types.Invoice, err error) {
	invoice, err = m.next.CalculateInvoice(obuID)
	return
}
