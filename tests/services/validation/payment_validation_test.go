package services_test

import (
	"flyhorizons-paymentservice/models"
	"flyhorizons-paymentservice/services/errors"
	"flyhorizons-paymentservice/services/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestPaymentService struct {
}

// Setup
func setupPaymentValidation() *validation.PaymentCredentialsValidation {
	paymentValidation := new(validation.PaymentCredentialsValidation)
	return paymentValidation
}

// Service Unit Tests
func TestProcessValidPaymentReturnsNilNoErrors(t *testing.T) {
	// Arrange
	paymentValidation := setupPaymentValidation()
	mockPayment := models.Payment{
		IBAN:      "DE44500105175407324931",
		CVV:       "123",
		FirstName: "Kate",
		LastName:  "Austen",
		Amount:    100,
		Currency:  "EUR",
	}

	// Act
	result := paymentValidation.CheckPaymentDetails(mockPayment)

	// Assert
	assert.Nil(t, result)
	assert.NoError(t, result)
}

func TestProcessPaymentInvalidIBANLengthReturnsError(t *testing.T) {
	// Arrange
	paymentValidation := setupPaymentValidation()
	mockPayment := models.Payment{
		IBAN:      "DE44500105",
		CVV:       "123",
		FirstName: "Kate",
		LastName:  "Austen",
		Amount:    100,
		Currency:  "EUR",
	}

	// Act
	result := paymentValidation.CheckPaymentDetails(mockPayment)

	// Assert
	assert.Error(t, result)
	assert.Equal(t, errors.NewInvalidIBANLengthError(400), result)
}

func TestProcessPaymentNoInitialLettersReturnsError(t *testing.T) {
	// Arrange
	paymentValidation := setupPaymentValidation()
	mockPayment := models.Payment{
		IBAN:      "44500105175407324931",
		CVV:       "123",
		FirstName: "Kate",
		LastName:  "Austen",
		Amount:    100,
		Currency:  "EUR",
	}

	// Act
	result := paymentValidation.CheckPaymentDetails(mockPayment)

	// Assert
	assert.Error(t, result)
	assert.Equal(t, errors.NewInvalidIBANNoLettersError(400), result)
}

func TestProcessPaymentInvalidPrefixReturnsError(t *testing.T) {
	// Arrange
	paymentValidation := setupPaymentValidation()
	mockPayment := models.Payment{
		IBAN:      "XX44500105175407324931",
		CVV:       "123",
		FirstName: "Kate",
		LastName:  "Austen",
		Amount:    100,
		Currency:  "EUR",
	}

	// Act
	result := paymentValidation.CheckPaymentDetails(mockPayment)

	// Assert
	assert.Error(t, result)
	assert.Equal(t, errors.NewInvalidIBANPrefix(400), result)
}

func TestProcessPaymentInvalidCVVLengthReturnsError(t *testing.T) {
	// Arrange
	paymentValidation := setupPaymentValidation()
	mockPayment := models.Payment{
		IBAN:      "DE44500105175407324931",
		CVV:       "12",
		FirstName: "Kate",
		LastName:  "Austen",
		Amount:    100,
		Currency:  "EUR",
	}

	// Act
	result := paymentValidation.CheckPaymentDetails(mockPayment)

	// Assert
	assert.Error(t, result)
	assert.Equal(t, errors.NewInvalidCVVLengthError(400), result)
}

func TestProcessPaymentInvalidCVVCharactersFormatReturnsError(t *testing.T) {
	// Arrange
	paymentValidation := setupPaymentValidation()
	mockPayment := models.Payment{
		IBAN:      "DE44500105175407324931",
		CVV:       "CBA",
		FirstName: "Kate",
		LastName:  "Austen",
		Amount:    100,
		Currency:  "EUR",
	}

	// Act
	result := paymentValidation.CheckPaymentDetails(mockPayment)

	// Assert
	assert.Error(t, result)
	assert.Equal(t, errors.NewInvalidCVVNoIntegersError(400), result)
}
