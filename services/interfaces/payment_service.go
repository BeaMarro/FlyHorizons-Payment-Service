package interfaces

import "flyhorizons-paymentservice/models"

type PaymentService interface {
	ProcessPayment(models.PaymentRequest) error
}
