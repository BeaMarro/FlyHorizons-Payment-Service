package services

import (
	"encoding/json"
	"flyhorizons-paymentservice/config"
	"flyhorizons-paymentservice/models"
	"flyhorizons-paymentservice/services/validation"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type PaymentService struct {
	paymentCredentialsValidation validation.PaymentCredentialsValidation
	paymentIntegrityValidation   validation.PaymentIntegrityValidation
	rabbitMQClient               *config.RabbitMQ
}

func NewPaymentService(paymentValidation validation.PaymentCredentialsValidation, client *config.RabbitMQ) *PaymentService {
	return &PaymentService{
		paymentCredentialsValidation: paymentValidation,
		rabbitMQClient:               client,
	}
}

func (paymentService *PaymentService) StartPaymentConsumer() {
	channel := paymentService.rabbitMQClient.Channel

	messages, err := channel.Consume(
		"booking.created", // Queue name
		"",                // Consumer tag
		true,              // Auto acknowledge
		false,             // Exclusive
		false,             // No-local
		false,             // No-wait
		nil,               // Arguments
	)
	if err != nil {
		log.Fatalf("An error occurred while registering the consumer: %v", err)
	}

	// Log consumer setup
	log.Printf("Started consumer for RabbitMQ queue: %s", "booking.created")
	log.Printf("Channel: RabbitMQ channel is active")

	// Start goroutine to process messages
	go func() {
		for message := range messages {
			// Log the message body
			log.Printf("Messages:")
			log.Printf("Body: %s", string(message.Body))

			var paymentRequest models.PaymentRequest
			// Convert from a JSON to PaymentRequest object
			if err := json.Unmarshal(message.Body, &paymentRequest); err != nil {
				log.Printf("An error occurred while converting the JSON to a PaymentRequest object: %v", err)
				continue
			}

			// Process the payment details once a new PaymentRequest has been posted
			paymentService.ProcessPayment(paymentRequest)
		}
		// Log when the consumer stops (e.g., if the channel closes)
		log.Printf("Consumer for queue %s stopped", "booking.created")
	}()

	log.Println("Waiting for a Booking to be created. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever // Keeps the program running
}

func (paymentService *PaymentService) ProcessPayment(paymentRequest models.PaymentRequest) error {
	// --- Validate the payment ---
	payment := paymentRequest.Payment
	// Validate the IBAN and CVV
	credentialsValidationResult := paymentService.paymentCredentialsValidation.CheckPaymentDetails(payment)
	// Validate the nonce and timestamp
	integrityValidationResult := paymentService.paymentIntegrityValidation.CheckPaymentIntegrity(payment)

	var event string
	if credentialsValidationResult == nil && integrityValidationResult == nil {
		event = "payment.success"
		log.Print("Payment succeeded")
	} else {
		event = "payment.failed"
		log.Print("Payment failed")
		log.Print("Credentials validation result:")
		log.Print(credentialsValidationResult)
		log.Print("Integrity validation result:")
		log.Print(integrityValidationResult)
	}

	// --- Marshal bookingID properly as JSON number ---
	body, err := json.Marshal(paymentRequest.BookingID)
	if err != nil {
		log.Printf("Error marshaling booking ID: %v", err)
		return err
	}

	// --- Publish ---
	channel := config.RabbitMQClient.Channel
	err = channel.Publish(
		"",
		event,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("Error publishing payment event to RabbitMQ: %v\n", err)
	}

	return credentialsValidationResult
}
