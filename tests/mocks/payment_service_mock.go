package mock_repositories

import (
	"flyhorizons-paymentservice/models"
	"flyhorizons-paymentservice/services/interfaces"

	"github.com/stretchr/testify/mock"
)

type MockPaymentService struct {
	mock.Mock
}

var _ interfaces.PaymentService = (*MockPaymentService)(nil)

func (m *MockPaymentService) ProcessPayment(paymentRequest models.PaymentRequest) error {
	args := m.Called(paymentRequest)
	return args.Error(0)
}
