package kafka

import (
	"encoding/json"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Kafka producer instance
var Producer *kafka.Producer
var topicName = "test_topic"

// MessagePayload defines the structure of the message
type MessagePayload struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// InitializeProducer initializes the Kafka producer
func InitializeProducer() error {
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}

	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		return err
	}
	return nil
}

// ProduceMessage sends a message to Kafka
func ProduceMessage(message MessagePayload) error {
	payloadBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
		Value:          payloadBytes,
	}

	err = Producer.Produce(kafkaMessage, nil)
	if err != nil {
		return err
	}
	// Wait for message deliveries
	Producer.Flush(15 * 1000)
	return nil
}
