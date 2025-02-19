package mdgeotrack

import "md-geo-track/request_response/location"

type Repository interface {
	// Health check
	HeartBeat() map[string]string
	ProcessLocation(req location.LocationReq) error
}
