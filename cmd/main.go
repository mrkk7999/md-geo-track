package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"md-geo-track/controller"
	"md-geo-track/implementation"
	"md-geo-track/kafka"
	"md-geo-track/repository"
	httpTransport "md-geo-track/transport/http"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// logger
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.Info("Starting application...")

	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.WithError(err).Fatal("Error loading .env file")
	}

	var (
		// Load configuration variables
		httpAddr = os.Getenv("HTTP_ADDR")

		// Database ENV variables
		dbHost     = os.Getenv("DB_HOST")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
		dbPort     = os.Getenv("DB_PORT")
		dbSSLMode  = os.Getenv("DB_SSLMODE")
		dbTimeZone = os.Getenv("DB_TIMEZONE")

		// Kafka ENV variables
		kafkaBrokers  = []string{os.Getenv("BROKERS")}
		kafkaTopic    = os.Getenv("TOPIC")
		maxRetries    = os.Getenv("MAX_RETRIES")
		retryInterval = os.Getenv("RETRY_INTERVAL")
	)

	// PostgreSQL DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone)

	// Connect to Database
	log.Info("Connecting to database...")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to database")
	}

	log.Info("Successfully connected to the database.")

	// Initialize repository
	repo := repository.New(db, log)

	// Ensure Kafka topic exists before creating the producer
	log.Info("Ensuring Kafka topic exists...")

	err = kafka.EnsureTopicExists(kafkaBrokers, kafkaTopic, log)
	if err != nil {
		log.WithError(err).Fatal("Failed to ensure Kafka topic exists")
	}

	// Convert Kafka configurations
	maxRetriesInt, err := strconv.Atoi(maxRetries)
	if err != nil {
		log.WithError(err).Fatal("Failed to convert MAX_RETRIES to int")
	}

	retryIntervalInt, err := strconv.Atoi(retryInterval)
	if err != nil {
		log.WithError(err).Fatal("Failed to convert RETRY_INTERVAL to int")
	}

	// Create Kafka producer
	kafkaConfig := kafka.NewKafkaConfig(kafkaBrokers, kafkaTopic, maxRetriesInt, time.Duration(retryIntervalInt)*time.Second)
	producer := kafka.NewSyncProducer(kafkaConfig, log)
	log.Info("Kafka producer initialized successfully.")

	// Initialize service and controller
	svc := implementation.New(repo, producer, kafkaTopic, log)
	controller := controller.New(svc, log)

	// Set up HTTP router
	handler := httpTransport.SetUpRouter(controller, log)

	// Handle OS signals for graceful shutdown
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.WithField("address", httpAddr).Info("Starting HTTP server...")

	go func() {
		server := &http.Server{
			Addr:    httpAddr,
			Handler: handler,
		}
		errs <- server.ListenAndServe()
	}()

	log.WithError(<-errs).Error("Application exiting due to error")
}
