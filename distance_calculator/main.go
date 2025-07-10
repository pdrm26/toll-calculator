package main

import (
	"log"
)

func main() {
	service := NewCalculateService()
	kafkaConsumer, err := NewkafkaConsumer("obudata", service)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()

}
