package middleware

import (
	logger "log"
	"net/http"

	"github.com/sirupsen/logrus"
)

// LoggingMiddleware logs the incoming request details and completion time
func LoggingMiddleware(next http.Handler, log *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println("#####################################################################################")

		log.Infof("Incoming request: Method=%s URL=%s", r.Method, r.URL.Path)

		// Call to next handler
		next.ServeHTTP(w, r)
		logger.Println("#####################################################################################")

	})
}
