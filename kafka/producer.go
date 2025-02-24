package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func PublishMessage(producer sarama.SyncProducer, topic, message string, log *logrus.Logger) error {
	// Create message to send to Kafka
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Send the message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.WithError(err).Error("Failed to send message")
		return err
	}

	log.WithFields(logrus.Fields{
		"topic":     topic,
		"partition": partition,
		"offset":    offset,
	}).Info("Message stored successfully")
	return nil
}
