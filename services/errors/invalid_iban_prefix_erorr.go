package errors

type InvalidIBANPrefix struct {
}

func (e *InvalidIBANPrefix) Error() string {
	return "The IBAN prefix is invalid, the prefix should not be 'XX', but a valid country code."
}

func NewInvalidIBANPrefix(errorCode int) *InvalidIBANPrefix {
	return &InvalidIBANPrefix{}
}
