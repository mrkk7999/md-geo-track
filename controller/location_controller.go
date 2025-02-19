package controller

import (
	"encoding/json"
	"md-geo-track/request_response/location"
	"net/http"
)

// ProcessLocation
func (c *Controller) ProcessLocation(w http.ResponseWriter, r *http.Request) {
	var req location.LocationReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := c.service.ProcessLocation(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := map[string]string{"message": "Location data saved successfully"}
	encodeJSONResponse(w, http.StatusOK, response, nil)
}
