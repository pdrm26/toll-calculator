package main

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pdrm26/toll-calculator/types"
)

type DataProducer interface {
	ProduceData(types.OBU) error
}

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(kafkaTopic string) (DataProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{
		producer: p,
		topic:    kafkaTopic,
	}, nil
}

func (k *KafkaProducer) ProduceData(data types.OBU) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.topic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)
}
