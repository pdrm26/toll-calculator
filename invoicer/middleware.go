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
