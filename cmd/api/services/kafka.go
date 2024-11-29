package services

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConfig holds the configuration for Kafka producer and consumer
type KafkaConfig struct {
	BootstrapServers string
	GroupID          string
	Topic            string
}

var Config = KafkaConfig{
	BootstrapServers: "kafka:9092",
	GroupID:          "tweet_group",
	Topic:            "tweet",
}

// NewProducer creates a new Kafka producer
func NewProducer(config KafkaConfig) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.BootstrapServers})
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// NewConsumer creates a new Kafka consumer
func NewConsumer(config KafkaConfig) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers,
		"group.id":          config.GroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// HandleProducerEvents handles the events from the producer
func HandleProducerEvents(producer *kafka.Producer) {
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Error sending message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Message sent to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}

// ProduceMessage sends a message to the specified topic
func ProduceMessage(producer *kafka.Producer, topic string, message []byte) error {
	return producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)
}

// ConsumeMessages consumes messages from the specified topic
func ConsumeMessages(consumer *kafka.Consumer, topic string) {
	err := consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	log.Println("Waiting for messages...")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
		} else {
			log.Printf("Message received: %s\n", string(msg.Value))
		}
	}
}

// CloseProducer closes the Kafka producer
func CloseProducer(producer *kafka.Producer) {
	producer.Close()
}

// CloseConsumer closes the Kafka consumer
func CloseConsumer(consumer *kafka.Consumer) {
	consumer.Close()
}
