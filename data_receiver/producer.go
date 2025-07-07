package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pdrm26/toll-calculator/types"
)

type DataProducer interface {
	ProduceData(types.OBU) error
}

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}
	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	return &KafkaProducer{
		producer: p,
	}, nil
}

func (k *KafkaProducer) ProduceData(data types.OBU) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)
}
