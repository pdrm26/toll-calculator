package main

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pdrm26/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	topic       string
	isRunning   bool // A signal handler or similar could be used to set this to false to break the loop.
	calcService CalculatorServicer
}

func NewkafkaConsumer(kafkaTopic string, service CalculatorServicer) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{kafkaTopic}, nil)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer:    c,
		topic:       kafkaTopic,
		calcService: service,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Kafka consumer started")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			log.Println("Kafka consumer error: ", err)
		}

		var obu types.OBU
		if err := json.Unmarshal(msg.Value, &obu); err != nil {
			logrus.Errorf("JSON serialization error: %s", err)
		}

		_, err = c.calcService.CalculateDistance(obu)
		if err != nil {
			logrus.Errorf("distance calculation error: %s", err)
		}

		// fmt.Printf("distance %.2f\n", distance)
	}
}
