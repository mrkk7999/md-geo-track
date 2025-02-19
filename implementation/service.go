package implementation

import (
	mdgeotrack "md-geo-track/iface"

	"github.com/IBM/sarama"
)

type service struct {
	repository mdgeotrack.Repository
	
	producer   sarama.SyncProducer
	topic      string
}

func New(repository mdgeotrack.Repository, producer sarama.SyncProducer, topic string) mdgeotrack.Service {
	return &service{
		repository: repository,
		producer:   producer,
		topic:      topic,
	}
}
