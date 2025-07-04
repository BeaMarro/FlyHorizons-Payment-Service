package errors

type ReplayAttackError struct {
}

func (e *ReplayAttackError) Error() string {
	return "The nonce has already been used, a nonce attack has been detected hence your request has been blocked by the server."
}

func NewReplayAttackError(errorCode int) *ReplayAttackError {
	return &ReplayAttackError{}
}
