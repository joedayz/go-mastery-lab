package main

import (
	"context"
	"fmt"
	"time"
)

// ============================================================================
// CONTEXT CON CANCELACIÓN
// ============================================================================
// Context es fundamental para manejar cancelación, timeouts y valores
// en operaciones concurrentes
// Similar a CompletableFuture.cancel() en Java, pero más integrado
// ============================================================================

// ============================================================================
// CONTEXT CON TIMEOUT
// ============================================================================

func contextWithTimeout() {
	fmt.Println("=== Context with Timeout ===")

	// Crear context con timeout de 2 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Simular operación que puede tardar
	done := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second) // Tarda más que el timeout
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Printf("Operation cancelled: %v\n", ctx.Err())
	}
	fmt.Println()
}

// ============================================================================
// CONTEXT CON CANCELACIÓN MANUAL
// ============================================================================

func contextWithCancel() {
	fmt.Println("=== Context with Cancel ===")

	ctx, cancel := context.WithCancel(context.Background())

	// Goroutine que hace trabajo
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker cancelled: %v\n", ctx.Err())
				return
			default:
				fmt.Printf("Working... %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// Cancelar después de 1 segundo
	time.Sleep(1 * time.Second)
	fmt.Println("Cancelling context...")
	cancel()
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
}

// ============================================================================
// CONTEXT CON DEADLINE
// ============================================================================
// Similar a timeout, pero con un tiempo absoluto

func contextWithDeadline() {
	fmt.Println("=== Context with Deadline ===")

	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Verificar deadline
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Printf("Deadline exceeded: %v\n", ctx.Err())
	}
	fmt.Println()
}

// ============================================================================
// CONTEXT CON VALORES
// ============================================================================
// Context puede llevar valores (útil para request IDs, user IDs, etc.)

func contextWithValues() {
	fmt.Println("=== Context with Values ===")

	// Crear context con valores
	ctx := context.WithValue(context.Background(), "userID", 123)
	ctx = context.WithValue(ctx, "requestID", "req-456")
	ctx = context.WithValue(ctx, "traceID", "trace-789")

	// Función que usa los valores
	processRequest := func(ctx context.Context) {
		userID := ctx.Value("userID").(int)
		requestID := ctx.Value("requestID").(string)
		traceID := ctx.Value("traceID").(string)

		fmt.Printf("Processing request:\n")
		fmt.Printf("  User ID: %d\n", userID)
		fmt.Printf("  Request ID: %s\n", requestID)
		fmt.Printf("  Trace ID: %s\n", traceID)
	}

	processRequest(ctx)
	fmt.Println()
}

// ============================================================================
// PROPAGAR CONTEXT EN FUNCIONES
// ============================================================================
// Siempre pasa context como primer parámetro en funciones que pueden cancelarse

func longRunningOperation(ctx context.Context, duration time.Duration) error {
	fmt.Printf("Starting operation (will take %v)...\n", duration)

	// Verificar si ya está cancelado
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Simular trabajo con checks periódicos de cancelación
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	start := time.Now()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Operation cancelled after %v\n", time.Since(start))
			return ctx.Err()
		case <-ticker.C:
			elapsed := time.Since(start)
			fmt.Printf("  Still working... (%v elapsed)\n", elapsed)
			if elapsed >= duration {
				fmt.Printf("Operation completed in %v\n", elapsed)
				return nil
			}
		}
	}
}

func propagateContext() {
	fmt.Println("=== Propagating Context ===")

	// Context con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Pasar context a función
	if err := longRunningOperation(ctx, 3*time.Second); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()
}

// ============================================================================
// CONTEXT EN HTTP REQUESTS
// ============================================================================
// Ejemplo de cómo usar context en operaciones HTTP

type HTTPClient struct{}

func (c *HTTPClient) Get(ctx context.Context, url string) (string, error) {
	// Simular request HTTP
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Second):
		return fmt.Sprintf("Response from %s", url), nil
	}
}

func contextInHTTP() {
	fmt.Println("=== Context in HTTP Requests ===")

	client := &HTTPClient{}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	result, err := client.Get(ctx, "https://api.example.com/data")
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	} else {
		fmt.Printf("Success: %s\n", result)
	}
	fmt.Println()
}

// ============================================================================
// CONTEXT EN DATABASE OPERATIONS
// ============================================================================
// Ejemplo de cómo usar context en operaciones de base de datos

type Database struct{}

func (db *Database) Query(ctx context.Context, query string) ([]string, error) {
	// Simular query que puede tardar
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(2 * time.Second):
		return []string{"result1", "result2", "result3"}, nil
	}
}

func contextInDatabase() {
	fmt.Println("=== Context in Database Operations ===")

	db := &Database{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	results, err := db.Query(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Printf("Query failed: %v\n", err)
	} else {
		fmt.Printf("Results: %v\n", results)
	}
	fmt.Println()
}

// ============================================================================
// CONTEXT CHAINING
// ============================================================================
// Puedes encadenar contexts para agregar más información o timeouts

func contextChaining() {
	fmt.Println("=== Context Chaining ===")

	// Context base con valores
	baseCtx := context.WithValue(context.Background(), "userID", 123)

	// Agregar timeout
	timeoutCtx, cancel1 := context.WithTimeout(baseCtx, 2*time.Second)
	defer cancel1()

	// Agregar más valores
	finalCtx := context.WithValue(timeoutCtx, "requestID", "req-999")

	// Usar el context final
	fmt.Printf("User ID: %d\n", finalCtx.Value("userID"))
	fmt.Printf("Request ID: %s\n", finalCtx.Value("requestID"))

	// El timeout también está presente
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed")
	case <-finalCtx.Done():
		fmt.Printf("Timeout: %v\n", finalCtx.Err())
	}
	fmt.Println()
}

// ============================================================================
// EJEMPLO PRÁCTICO: SERVICIO CON MÚLTIPLES OPERACIONES
// ============================================================================

type Service struct {
	db     *Database
	client *HTTPClient
}

func NewService() *Service {
	return &Service{
		db:     &Database{},
		client: &HTTPClient{},
	}
}

func (s *Service) ProcessData(ctx context.Context, userID int) error {
	// Agregar userID al context
	ctx = context.WithValue(ctx, "userID", userID)

	// Hacer múltiples operaciones con el mismo context
	// Si alguna falla o se cancela, todas se cancelan

	// 1. Query database
	results, err := s.db.Query(ctx, "SELECT * FROM data")
	if err != nil {
		return fmt.Errorf("database query failed: %w", err)
	}

	// 2. Fetch external data
	externalData, err := s.client.Get(ctx, "https://api.example.com/data")
	if err != nil {
		return fmt.Errorf("external API call failed: %w", err)
	}

	fmt.Printf("Processed data: %v, External: %s\n", results, externalData)
	return nil
}

func practicalExample() {
	fmt.Println("=== Practical Example: Service with Multiple Operations ===")

	service := NewService()

	// Context con timeout para toda la operación
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := service.ProcessData(ctx, 123); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	fmt.Println("=== CONCURRENCIA: CONTEXT ===\n")

	contextWithTimeout()
	contextWithCancel()
	contextWithDeadline()
	contextWithValues()
	propagateContext()
	contextInHTTP()
	contextInDatabase()
	contextChaining()
	practicalExample()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

