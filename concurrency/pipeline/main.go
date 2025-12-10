package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// ============================================================================
// PIPELINES
// ============================================================================
// Pipelines procesan datos en etapas usando channels
// Cada etapa es una goroutine que recibe de un channel y envía a otro
// Similar a Stream API en Java, pero más explícito y concurrente
// ============================================================================

// ============================================================================
// PIPELINE BÁSICO
// ============================================================================
// Pipeline simple: Generar -> Cuadrar -> Imprimir

func basicPipeline() {
	fmt.Println("=== Basic Pipeline ===")

	// Stage 1: Generar números
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
	}()

	// Stage 2: Cuadrar números
	squares := make(chan int)
	go func() {
		defer close(squares)
		for n := range numbers {
			squares <- n * n
		}
	}()

	// Stage 3: Imprimir resultados
	for square := range squares {
		fmt.Printf("Square: %d\n", square)
	}
	fmt.Println()
}

// ============================================================================
// PIPELINE CON MÚLTIPLES ETAPAS
// ============================================================================

func multiStagePipeline() {
	fmt.Println("=== Multi-Stage Pipeline ===")

	// Stage 1: Generar números
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 10; i++ {
			numbers <- i
		}
	}()

	// Stage 2: Filtrar pares
	evens := make(chan int)
	go func() {
		defer close(evens)
		for n := range numbers {
			if n%2 == 0 {
				evens <- n
			}
		}
	}()

	// Stage 3: Multiplicar por 10
	multiplied := make(chan int)
	go func() {
		defer close(multiplied)
		for n := range evens {
			multiplied <- n * 10
		}
	}()

	// Stage 4: Imprimir
	for result := range multiplied {
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Println()
}

// ============================================================================
// PIPELINE CON FUNCIONES REUTILIZABLES
// ============================================================================

// Función genérica para pipeline stage
func pipelineStage[T any](input <-chan T, output chan<- T, fn func(T) T) {
	defer close(output)
	for item := range input {
		output <- fn(item)
	}
}

func reusablePipeline() {
	fmt.Println("=== Reusable Pipeline ===")

	// Generar
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
	}()

	// Cuadrar
	squares := make(chan int)
	go pipelineStage(numbers, squares, func(n int) int {
		return n * n
	})

	// Imprimir
	for square := range squares {
		fmt.Printf("Square: %d\n", square)
	}
	fmt.Println()
}

// ============================================================================
// FAN-OUT / FAN-IN PATTERN
// ============================================================================
// Fan-out: Distribuir trabajo a múltiples workers
// Fan-in: Combinar resultados de múltiples workers

func fanOutFanIn() {
	fmt.Println("=== Fan-Out / Fan-In ===")

	// Input
	input := make(chan int)
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()

	// Fan-out: Múltiples workers procesan
	numWorkers := 3
	workerOutputs := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		output := make(chan int)
		workerOutputs[i] = output
		go func(workerID int) {
			defer close(output)
			for n := range input {
				// Simular trabajo
				time.Sleep(100 * time.Millisecond)
				result := n * n
				fmt.Printf("Worker %d processed %d -> %d\n", workerID, n, result)
				output <- result
			}
		}(i)
	}

	// Fan-in: Combinar resultados
	results := make(chan int)
	var wg sync.WaitGroup
	for _, workerOutput := range workerOutputs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for result := range ch {
				results <- result
			}
		}(workerOutput)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Recibir resultados combinados
	for result := range results {
		fmt.Printf("Final result: %d\n", result)
	}
	fmt.Println()
}

// ============================================================================
// PIPELINE CON RATE LIMITING
// ============================================================================

func rateLimitedPipeline() {
	fmt.Println("=== Rate Limited Pipeline ===")

	// Generar números rápidamente
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 10; i++ {
			numbers <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Rate limiter: Solo procesar uno cada 200ms
	rateLimited := make(chan int)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		defer close(rateLimited)
		for {
			select {
			case n, ok := <-numbers:
				if !ok {
					return
				}
				<-ticker.C // Esperar al siguiente tick
				rateLimited <- n
			}
		}
	}()

	// Procesar
	for n := range rateLimited {
		fmt.Printf("Processed: %d\n", n)
	}
	fmt.Println()
}

// ============================================================================
// EJEMPLO PRÁCTICO: PIPELINE DE TRANSFORMACIÓN DE DATOS
// ============================================================================
// Este es el ejemplo sugerido en el temario

type DataPoint struct {
	Value float64
	Time  time.Time
}

func dataTransformationPipeline() {
	fmt.Println("=== Data Transformation Pipeline ===")

	// Stage 1: Generar datos
	rawData := make(chan DataPoint)
	go func() {
		defer close(rawData)
		baseTime := time.Now()
		for i := 0; i < 10; i++ {
			rawData <- DataPoint{
				Value: float64(i),
				Time:  baseTime.Add(time.Duration(i) * time.Second),
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Stage 2: Filtrar valores válidos (> 0)
	filtered := make(chan DataPoint)
	go func() {
		defer close(filtered)
		for point := range rawData {
			if point.Value > 0 {
				filtered <- point
			}
		}
	}()

	// Stage 3: Transformar (aplicar función matemática)
	transformed := make(chan DataPoint)
	go func() {
		defer close(transformed)
		for point := range filtered {
			transformed <- DataPoint{
				Value: math.Sqrt(point.Value) * 10, // Transformación
				Time:  point.Time,
			}
		}
	}()

	// Stage 4: Agregar metadata
	enriched := make(chan DataPoint)
	go func() {
		defer close(enriched)
		for point := range transformed {
			// Enriquecer con metadata (en este caso solo imprimir)
			fmt.Printf("Enriched: Value=%.2f, Time=%s\n", point.Value, point.Time.Format("15:04:05"))
			enriched <- point
		}
	}()

	// Consumir resultados finales
	count := 0
	for range enriched {
		count++
	}
	fmt.Printf("Processed %d data points\n", count)
	fmt.Println()
}

// ============================================================================
// PIPELINE CON ERROR HANDLING
// ============================================================================

type Result struct {
	Value int
	Error error
}

func pipelineWithErrorHandling() {
	fmt.Println("=== Pipeline with Error Handling ===")

	// Generar números (algunos inválidos)
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := -2; i <= 5; i++ {
			numbers <- i
		}
	}()

	// Procesar con validación
	results := make(chan Result)
	go func() {
		defer close(results)
		for n := range numbers {
			if n < 0 {
				results <- Result{Error: fmt.Errorf("negative number: %d", n)}
				continue
			}
			results <- Result{Value: n * n}
		}
	}()

	// Manejar resultados
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
		} else {
			fmt.Printf("Success: %d\n", result.Value)
		}
	}
	fmt.Println()
}

// ============================================================================
// PIPELINE CON BUFFERING
// ============================================================================
// Buffers pueden mejorar el throughput al permitir que stages trabajen
// en paralelo sin bloquearse

func bufferedPipeline() {
	fmt.Println("=== Buffered Pipeline ===")

	// Buffers permiten que stages trabajen en paralelo
	stage1 := make(chan int, 5)
	stage2 := make(chan int, 5)
	stage3 := make(chan int, 5)

	// Stage 1: Generar
	go func() {
		defer close(stage1)
		for i := 1; i <= 10; i++ {
			stage1 <- i
			fmt.Printf("Generated: %d\n", i)
		}
	}()

	// Stage 2: Procesar
	go func() {
		defer close(stage2)
		for n := range stage1 {
			time.Sleep(100 * time.Millisecond) // Trabajo lento
			stage2 <- n * 2
			fmt.Printf("Processed: %d -> %d\n", n, n*2)
		}
	}()

	// Stage 3: Consumir
	for result := range stage2 {
		fmt.Printf("Final: %d\n", result)
	}
	fmt.Println()
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	fmt.Println("=== CONCURRENCIA: PIPELINES ===\n")

	basicPipeline()
	multiStagePipeline()
	reusablePipeline()
	fanOutFanIn()
	rateLimitedPipeline()
	dataTransformationPipeline()
	pipelineWithErrorHandling()
	bufferedPipeline()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

