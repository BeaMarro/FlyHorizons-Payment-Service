package errors

type InvalidIBANLengthError struct {
}

func (e *InvalidIBANLengthError) Error() string {
	return "The IBAN has an invalid length, it should be between 15 and 34 characters inclusive to ensure validity."
}

func NewInvalidIBANLengthError(errorCode int) *InvalidIBANLengthError {
	return &InvalidIBANLengthError{}
}
