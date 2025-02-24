package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

// EnsureTopicExists checks if a Kafka topic exists and creates it if necessary.
func EnsureTopicExists(brokers []string, topic string, log *logrus.Logger) error {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0 // Matches your Kafka setup
	config.Producer.Return.Successes = true

	// Create a new Kafka admin client
	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		log.Errorf("Error creating Kafka admin client: %v", err)
		return err
	}
	defer admin.Close()

	// Check if the topic already exists
	topics, err := admin.ListTopics()
	if err != nil {
		log.Errorf("Error listing Kafka topics: %v", err)
		return err
	}

	if _, exists := topics[topic]; exists {
		log.Infof("✅ Topic '%s' already exists, skipping creation.", topic)
		return nil
	}

	// Define topic configuration
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     1, // Adjust if needed
		ReplicationFactor: 1, // Should match KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR
	}

	// Create topic
	err = admin.CreateTopic(topic, topicDetail, false)
	if err != nil {
		log.Errorf("⚠️ Failed to create topic '%s': %v", topic, err)
		return err
	}

	log.Infof("✅ Topic '%s' created successfully.", topic)
	return nil
}
