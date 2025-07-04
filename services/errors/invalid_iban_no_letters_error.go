package errors

type InvalidIBANNoLettersError struct {
}

func (e *InvalidIBANNoLettersError) Error() string {
	return "The IBAN is invalid, it should begin with 2 letters."
}

func NewInvalidIBANNoLettersError(errorCode int) *InvalidIBANNoLettersError {
	return &InvalidIBANNoLettersError{}
}
