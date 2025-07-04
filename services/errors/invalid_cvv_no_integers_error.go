package errors

type InvalidCVVNoIntegersError struct {
}

func (e *InvalidCVVNoIntegersError) Error() string {
	return "The CVV is invalid, it should have integer characters, not letters."
}

func NewInvalidCVVNoIntegersError(errorCode int) *InvalidCVVNoIntegersError {
	return &InvalidCVVNoIntegersError{}
}
