package main

import (
	"log"
)

func main() {
	service := NewCalculatorService()
	service = NewLogMiddleware(service)
	kafkaConsumer, err := NewkafkaConsumer("obudata", service)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()

}
