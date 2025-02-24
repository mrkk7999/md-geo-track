package controller

import (
	"encoding/json"
	"errors"
	"md-geo-track/request_response/location"
	"net/http"
)

// ProcessLocation
func (c *Controller) ProcessLocation(w http.ResponseWriter, r *http.Request) {

	var req location.LocationReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.log.WithError(err).Error("Invalid request payload")
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("invalid request payload"))
		return
	}

	err := c.service.ProcessLocation(req)
	if err != nil {
		c.log.WithError(err).Error("Error processing location data")
		encodeJSONResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	response := map[string]string{"message": "Location data saved successfully"}
	encodeJSONResponse(w, http.StatusOK, response, nil)
}
