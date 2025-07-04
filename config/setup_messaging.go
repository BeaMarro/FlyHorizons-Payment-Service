package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

type SetupMessaging struct {
}

func InitializeRabbitMQ() *RabbitMQ {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	connection, err := amqp091.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("An error occurred while connecting to RabbitMQ: %s", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("An error occurred while opening the RabbitMQ channel: %s", err)
	}

	// Declare the booking.created queue
	_, err = channel.QueueDeclare(
		"booking.created",
		true,  // Durable
		false, // Auto Delete
		false, // Exclusive
		false, // No Wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatalf("An error occurred while declaring the booking.created queue: %s", err)
	}

	// Declare the payment.failed queue
	_, err = channel.QueueDeclare(
		"payment.failed",
		true,  // Durable
		false, // Auto Delete
		false, // Exclusive
		false, // No Wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatalf("An error occurred while declaring the payment.failed queue: %s", err)
	}

	// Declare the payment.success queue
	_, err = channel.QueueDeclare(
		"payment.success",
		true,  // Durable
		false, // Auto Delete
		false, // Exclusive
		false, // No Wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatalf("An error occurred while declaring the payment.success queue: %s", err)
	}

	log.Println("RabbitMQ has been initialized successfully.")

	return &RabbitMQ{
		Connection: connection,
		Channel:    channel,
	}
}
