package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

// ============================================================================
// RETRY CON BACKOFF
// ============================================================================
// Patrón común para reintentar operaciones que pueden fallar
// Backoff exponencial reduce la carga en el sistema
// ============================================================================

var ErrTemporaryFailure = errors.New("temporary failure")

// RetryWithBackoff reintenta una operación con backoff exponencial
func RetryWithBackoff(maxRetries int, initialDelay time.Duration, fn func() error) error {
	var err error
	delay := initialDelay

	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}

		// Verificar si es un error temporal
		if !errors.Is(err, ErrTemporaryFailure) {
			return err // Error permanente, no reintentar
		}

		if i < maxRetries-1 {
			fmt.Printf("Retry %d/%d after %v: %v\n", i+1, maxRetries, delay, err)
			time.Sleep(delay)
			delay = time.Duration(float64(delay) * 1.5) // Backoff exponencial
		}
	}

	return fmt.Errorf("failed after %d retries: %w", maxRetries, err)
}

// ExponentialBackoff calcula delay exponencial
func ExponentialBackoff(attempt int, baseDelay time.Duration) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * baseDelay
}

func main() {
	attempt := 0
	operation := func() error {
		attempt++
		fmt.Printf("Attempt %d\n", attempt)
		if attempt < 3 {
			return ErrTemporaryFailure
		}
		return nil
	}

	err := RetryWithBackoff(5, 100*time.Millisecond, operation)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Success!")
	}
}

