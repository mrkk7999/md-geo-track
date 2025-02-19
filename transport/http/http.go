package http

import (
	"md-geo-track/controller"
	"md-geo-track/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRouter(controller *controller.Controller) http.Handler {
	var (
		router = mux.NewRouter()
		// subRouter = router.PathPrefix("/loc/api/v1").Subrouter()
	)
	// Logging middleware
	router.Use(middleware.LoggingMiddleware)
	// Healthcheck
	router.HandleFunc("/loc/api/v1/heartbeat", controller.HeartBeatHandler).Methods("GET")
	// Location data submission API
	router.HandleFunc("/loc/api/v1/submit", controller.ProcessLocation).Methods("POST")

	return router
}
