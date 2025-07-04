package main

import (
	"flyhorizons-paymentservice/config"
	"flyhorizons-paymentservice/internal/health"
	"flyhorizons-paymentservice/internal/metrics"
	"flyhorizons-paymentservice/services"
	"flyhorizons-paymentservice/services/validation"

	"github.com/gin-gonic/gin"

	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	healthcfg "github.com/tavsec/gin-healthcheck/config"
)

func main() {
	// Initialize RabbitMQ for messaging
	rabbitMQClient := config.InitializeRabbitMQ()
	config.RabbitMQClient = rabbitMQClient

	defer rabbitMQClient.Channel.Close()
	defer rabbitMQClient.Connection.Close()

	router := gin.Default()

	// --- Health checks setup ---
	conf := healthcfg.DefaultConfig()
	rabbitMQCheck := health.RabbitMQCheck{}
	healthcheck.New(router, conf, []checks.Check{rabbitMQCheck})

	// --- Metrics setup ---
	metrics.RegisterMetricsRoutes(router, rabbitMQCheck)

	// --- Microservice setup ---
	paymentValidation := validation.PaymentCredentialsValidation{}
	paymentService := services.NewPaymentService(paymentValidation, rabbitMQClient)
	go paymentService.StartPaymentConsumer()

	// Run the microservice
	router.Run(":8084")
}
