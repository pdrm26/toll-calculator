package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer  *kafka.Consumer
	topic     string
	isRunning bool // A signal handler or similar could be used to set this to false to break the loop.
}

func NewkafkaConsumer(kafkaTopic string) (*KafkaConsumer, error) {
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
		consumer: c,
		topic:    kafkaTopic,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Kafka consumer started")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(time.Second)
		if err != nil {
			log.Println("Kafka consumer error: ", err)
		}
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}

		fmt.Println(msg)
	}
}
