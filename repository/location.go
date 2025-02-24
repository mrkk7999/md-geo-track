package repository

import (
	"md-geo-track/request_response/location"
	"time"
)

func (r *repository) ProcessLocation(req location.LocationReq) error {

	// Insert location data
	loc := location.LocationModel{
		TenantID:  req.TenantID,
		UserID:    req.UserID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Timestamp: time.Now(),
	}

	if err := r.db.Create(&loc).Error; err != nil {
		return err // Return error directly, no logging here
	}

	r.log.Info("Location data saved successfully")

	return nil
}
