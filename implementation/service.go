package implementation

import (
	mdgeotrack "md-geo-track/iface"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type service struct {
	repository mdgeotrack.Repository

	producer sarama.SyncProducer
	topic    string

	log *logrus.Logger
}

func New(repository mdgeotrack.Repository, producer sarama.SyncProducer, topic string, log *logrus.Logger) mdgeotrack.Service {
	return &service{
		repository: repository,
		producer:   producer,
		topic:      topic,
		log:        log,
	}
}
