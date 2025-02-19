package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func PublishMessage(producer sarama.SyncProducer, topic, message string) error {
	// Create message to send to Kafka
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Send the message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	log.Printf("Message stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
