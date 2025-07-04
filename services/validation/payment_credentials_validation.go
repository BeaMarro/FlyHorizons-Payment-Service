package validation

import (
	"flyhorizons-paymentservice/models"
	"flyhorizons-paymentservice/services/errors"
	"strings"
	"unicode"
)

type PaymentCredentialsValidation struct{}

func (paymentCredentialsValidation *PaymentCredentialsValidation) checkIBANFormat(iban string) error {
	// IBAN length needs to be between 15 and 34 characters inclusive
	if len(iban) < 15 || len(iban) > 34 {
		return errors.NewInvalidIBANLengthError(400)
	}
	// IBAN should start with 2 letters
	if !unicode.IsLetter(rune(iban[0])) || !unicode.IsLetter(rune(iban[1])) {
		return errors.NewInvalidIBANNoLettersError(400)
	}
	// IBANs cannot begin with characters "XX"
	if strings.HasPrefix(iban, "XX") {
		return errors.NewInvalidIBANPrefix(400)
	}
	return nil
}

func (pv *PaymentCredentialsValidation) checkCVV(cvv string) error {
	// CVV needs to consist of 3 numeric digits
	if len(cvv) != 3 {
		return errors.NewInvalidCVVLengthError(400)
	}
	for _, r := range cvv {
		if !unicode.IsDigit(r) {
			return errors.NewInvalidCVVNoIntegersError(400)
		}
	}
	return nil
}

func (pv *PaymentCredentialsValidation) CheckPaymentDetails(payment models.Payment) error {
	if err := pv.checkIBANFormat(payment.IBAN); err != nil {
		return err
	}
	if err := pv.checkCVV(payment.CVV); err != nil {
		return err
	}
	return nil // success
}
