package errors

type ExpiredTimestampError struct {
}

func (e *ExpiredTimestampError) Error() string {
	return "The payment timestamp has expired, therefore the request has been rejected by the server."
}

func NewExpiredTimestampError(errorCode int) *ExpiredTimestampError {
	return &ExpiredTimestampError{}
}
