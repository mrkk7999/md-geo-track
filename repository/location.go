package repository

import (
	"errors"
	"log"
	"md-geo-track/request_response/location"
	"time"
)

func (r *repository) ProcessLocation(req location.LocationReq) error {
	tx := r.db.Begin()

	// Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Insert location data
	loc := location.LocationModel{
		TenantID:  req.TenantID,
		UserID:    req.UserID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Timestamp: time.Now(),
	}

	if err := tx.Create(&loc).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to insert location data")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("failed to commit transaction")
	}

	log.Println("Location data saved successfully for TenantID:", req.TenantID, "UserID:", req.UserID)

	return nil
}
