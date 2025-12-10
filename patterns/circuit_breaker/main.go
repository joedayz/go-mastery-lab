package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ============================================================================
// CIRCUIT BREAKER PATTERN
// ============================================================================
// Previene llamadas a servicios que están fallando
// Similar a Hystrix en Java
// ============================================================================

type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

type CircuitBreaker struct {
	mu                sync.RWMutex
	state             State
	failureCount      int
	successCount      int
	maxFailures       int
	resetTimeout      time.Duration
	lastFailureTime   time.Time
	halfOpenMaxCalls  int
}

func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            StateClosed,
		maxFailures:      maxFailures,
		resetTimeout:     resetTimeout,
		halfOpenMaxCalls: 3,
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	// Verificar si debemos cambiar de estado
	cb.checkState()

	switch cb.state {
	case StateOpen:
		return errors.New("circuit breaker is open")
	case StateHalfOpen:
		if cb.successCount >= cb.halfOpenMaxCalls {
			cb.state = StateClosed
			cb.successCount = 0
		}
	}

	// Ejecutar función
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.onFailure()
		return err
	}

	cb.onSuccess()
	return nil
}

func (cb *CircuitBreaker) checkState() {
	if cb.state == StateOpen {
		if time.Since(cb.lastFailureTime) >= cb.resetTimeout {
			cb.state = StateHalfOpen
			cb.successCount = 0
		}
	}
}

func (cb *CircuitBreaker) onFailure() {
	cb.failureCount++
	cb.lastFailureTime = time.Now()

	if cb.failureCount >= cb.maxFailures {
		cb.state = StateOpen
	}
}

func (cb *CircuitBreaker) onSuccess() {
	cb.failureCount = 0
	if cb.state == StateHalfOpen {
		cb.successCount++
	}
}

func main() {
	cb := NewCircuitBreaker(3, 2*time.Second)

	// Simular operaciones que fallan
	for i := 0; i < 5; i++ {
		err := cb.Call(func() error {
			return errors.New("service error")
		})
		if err != nil {
			fmt.Printf("Call %d failed: %v\n", i+1, err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	// Esperar reset
	fmt.Println("Waiting for circuit breaker to reset...")
	time.Sleep(3 * time.Second)

	// Intentar de nuevo
	err := cb.Call(func() error {
		fmt.Println("Service call succeeded")
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

