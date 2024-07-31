package main

import (
	"log"

	"github.com/your-username/kafka-go-app/internal/kafka"
)

func main() {
	// Initialize Kafka Producer
	err := kafka.InitializeProducer()
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}

	// Create a test message
	message := kafka.MessagePayload{
		ID:      "1",
		Content: "Hello Kafka!",
	}

	// Produce the message to Kafka
	err = kafka.ProduceMessage(message)
	if err != nil {
		log.Fatalf("Failed to produce message: %v", err)
	}

	log.Println("Message successfully sent to Kafka")
}
