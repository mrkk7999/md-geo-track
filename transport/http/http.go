package http

import (
	"md-geo-track/controller"
	"md-geo-track/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func SetUpRouter(controller *controller.Controller, log *logrus.Logger) http.Handler {
	var (
		router = mux.NewRouter()
		// subRouter = router.PathPrefix("/loc/api/v1").Subrouter()
	)
	// Logging middleware
	router.Use(func(next http.Handler) http.Handler {
		return middleware.LoggingMiddleware(next, log)
	})
	// Healthcheck
	router.HandleFunc("/loc/api/v1/heartbeat", controller.HeartBeatHandler).Methods("GET")
	// Location data submission API
	router.HandleFunc("/loc/api/v1/submit", controller.ProcessLocation).Methods("POST")

	return router
}
