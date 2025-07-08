package main

import (
	"time"

	"github.com/pdrm26/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data types.OBU) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.ID,
			"lat":   data.Lat,
			"long":  data.Long,
			"took":  time.Since(start),
		}).Info("producing to kafka")
	}(time.Now())

	return l.next.ProduceData(data)
}
