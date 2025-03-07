package controller

import (
	"encoding/json"
	mdgeotrack "md-geo-track/iface"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Controller struct {
	service mdgeotrack.Service
	log     *logrus.Logger
}

func New(service mdgeotrack.Service, log *logrus.Logger) *Controller {
	return &Controller{
		service: service,
		log:     log,
	}
}

func encodeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
