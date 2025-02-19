package implementation

import (
	"encoding/json"
	"md-geo-track/kafka"
	"md-geo-track/request_response/location"
)

func (s *service) ProcessLocation(req location.LocationReq) error {
	// Store location data in the repository
	err := s.repository.ProcessLocation(req)
	if err != nil {
		return err
	}

	message, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// Publish location data to Kafka with retries
	return kafka.PublishMessage(s.producer, s.topic, string(message))
}
