package kafka

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

type KafkaConfig struct {
	Brokers       []string
	Topic         string
	MaxRetries    int
	RetryInterval time.Duration
}

func NewKafkaConfig(brokers []string, topic string, maxRetries int, retryInterval time.Duration) *KafkaConfig {
	return &KafkaConfig{
		Brokers:       brokers,
		Topic:         topic,
		MaxRetries:    maxRetries,
		RetryInterval: retryInterval,
	}
}

func NewSyncProducer(config *KafkaConfig) sarama.SyncProducer {
	producerConfig := sarama.NewConfig()
	
	producerConfig.Producer.Retry.Max = config.MaxRetries
	producerConfig.Producer.Retry.Backoff = config.RetryInterval
	producerConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(config.Brokers, producerConfig)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
	return producer
}
