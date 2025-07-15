package main

import (
	"time"

	"github.com/pdrm26/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (l LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  start,
			"error": err,
		}).Info("AggregateDistance")
	}(time.Now())

	err = l.next.AggregateDistance(distance)
	return
}

func (l LogMiddleware) CalculateInvoice(obuID int) (invoice *types.Invoice, err error) {
	var (
		totalDistance float64
		totalPrice    float64
	)
	if invoice != nil {
		totalDistance = invoice.TotalDistance
		totalPrice = invoice.TotalPrice
	}
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":          start,
			"error":         err,
			"totalDistance": totalDistance,
			"totalPrice":    totalPrice,
			"OBUID":         obuID,
		})
	}(time.Now())

	invoice, err = l.next.CalculateInvoice(obuID)
	return

}
