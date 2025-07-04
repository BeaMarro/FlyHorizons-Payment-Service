package validation

import (
	"flyhorizons-paymentservice/models"
	"flyhorizons-paymentservice/services/errors"
	"log"
	"sync"
	"time"
)

type PaymentIntegrityValidation struct {
	validatedNonces map[string]time.Time
	mu              sync.Mutex
}

// Constructor that initializes the nonce map
func NewPaymentIntegrityValidation() *PaymentIntegrityValidation {
	return &PaymentIntegrityValidation{
		validatedNonces: make(map[string]time.Time),
	}
}

// Ensures the payment timestamp is within ±5 minutes of the current UTC time
func (pv *PaymentIntegrityValidation) checkTimestamp(timestamp time.Time) error {
	now := time.Now().UTC()
	diff := now.Sub(timestamp)

	log.Printf("Now (UTC): %s", now.Format(time.RFC3339))
	log.Printf("Incoming timestamp: %s", timestamp.Format(time.RFC3339))
	log.Printf("Difference in seconds: %f", diff.Seconds())

	if diff < -10*time.Minute || diff > 10*time.Minute {
		return errors.NewExpiredTimestampError(400)
	}
	return nil
}

// Ensures the nonce is unique for the session
func (pv *PaymentIntegrityValidation) checkNonce(nonce string) error {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	// Defensive guard in case someone skips the constructor
	if pv.validatedNonces == nil {
		log.Println("[WARNING] validatedNonces map was nil — initializing defensively")
		pv.validatedNonces = make(map[string]time.Time)
	}

	if _, exists := pv.validatedNonces[nonce]; exists {
		return errors.NewReplayAttackError(400)
	}

	pv.validatedNonces[nonce] = time.Now()
	return nil
}

// Main public method
func (pv *PaymentIntegrityValidation) CheckPaymentIntegrity(payment models.Payment) error {
	if err := pv.checkTimestamp(payment.Timestamp); err != nil {
		return err
	}
	if err := pv.checkNonce(payment.Nonce); err != nil {
		return err
	}
	return nil
}
