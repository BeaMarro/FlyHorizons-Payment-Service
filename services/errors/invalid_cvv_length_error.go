package errors

type InvalidCVVLengthError struct {
}

func (e *InvalidCVVLengthError) Error() string {
	return "The CVV has an invalid length, it should be 3 characters to ensure validity."
}

func NewInvalidCVVLengthError(errorCode int) *InvalidCVVLengthError {
	return &InvalidCVVLengthError{}
}
