package kafka

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
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

func NewSyncProducer(config *KafkaConfig, log *logrus.Logger) sarama.SyncProducer {
	producerConfig := sarama.NewConfig()

	producerConfig.Producer.Retry.Max = config.MaxRetries
	producerConfig.Producer.Retry.Backoff = config.RetryInterval
	producerConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(config.Brokers, producerConfig)
	if err != nil {
		log.Error("Error creating Kafka producer: ", err)
		return nil
	}
	return producer
}
