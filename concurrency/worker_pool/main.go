package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ============================================================================
// WORKER POOLS
// ============================================================================
// Worker pools son un patrón común para procesar múltiples tareas
// con un número limitado de workers concurrentes
// Similar a ThreadPoolExecutor en Java
// ============================================================================

// ============================================================================
// WORKER POOL BÁSICO
// ============================================================================

type Task struct {
	ID   int
	Data string
}

type Result struct {
	TaskID int
	Output string
	Error  error
}

func basicWorkerPool() {
	fmt.Println("=== Basic Worker Pool ===")

	// Configuración
	numWorkers := 3
	numTasks := 10

	// Channels
	jobs := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	// Crear workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range jobs {
				// Simular procesamiento
				time.Sleep(200 * time.Millisecond)
				output := fmt.Sprintf("Processed by worker %d: %s", workerID, task.Data)
				results <- Result{
					TaskID: task.ID,
					Output: output,
					Error:  nil,
				}
			}
		}(w)
	}

	// Enviar tareas
	for i := 1; i <= numTasks; i++ {
		jobs <- Task{
			ID:   i,
			Data: fmt.Sprintf("Task %d", i),
		}
	}
	close(jobs)

	// Cerrar results cuando todos los workers terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recibir resultados
	for result := range results {
		fmt.Printf("Result: %s\n", result.Output)
	}
	fmt.Println()
}

// ============================================================================
// WORKER POOL CON CONTEXT (CANCELACIÓN)
// ============================================================================

func workerPoolWithContext() {
	fmt.Println("=== Worker Pool with Context ===")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	numWorkers := 3
	jobs := make(chan Task, 10)
	results := make(chan Result, 10)

	// Workers con context
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d: Cancelled\n", workerID)
					return
				case task, ok := <-jobs:
					if !ok {
						return
					}
					// Procesar tarea
					time.Sleep(500 * time.Millisecond)
					select {
					case results <- Result{TaskID: task.ID, Output: fmt.Sprintf("Processed %s", task.Data)}:
					case <-ctx.Done():
						return
					}
				}
			}
		}(w)
	}

	// Enviar tareas
	go func() {
		defer close(jobs)
		for i := 1; i <= 10; i++ {
			select {
			case jobs <- Task{ID: i, Data: fmt.Sprintf("Task %d", i)}:
			case <-ctx.Done():
				return
			}
		}
	}()

	// Cerrar results cuando terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recibir resultados
	for result := range results {
		fmt.Printf("Result: %s\n", result.Output)
	}
	fmt.Println()
}

// ============================================================================
// WORKER POOL REUTILIZABLE
// ============================================================================

type WorkerPool struct {
	numWorkers int
	jobs       chan Task
	results    chan Result
	wg         sync.WaitGroup
}

func NewWorkerPool(numWorkers, jobBufferSize int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Task, jobBufferSize),
		results:    make(chan Result, jobBufferSize),
	}
}

func (wp *WorkerPool) Start(processor func(Task) Result) {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go func(workerID int) {
			defer wp.wg.Done()
			for task := range wp.jobs {
				result := processor(task)
				result.TaskID = task.ID
				wp.results <- result
			}
		}(i)
	}
}

func (wp *WorkerPool) Submit(task Task) {
	wp.jobs <- task
}

func (wp *WorkerPool) Close() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
}

func (wp *WorkerPool) Results() <-chan Result {
	return wp.results
}

func reusableWorkerPool() {
	fmt.Println("=== Reusable Worker Pool ===")

	pool := NewWorkerPool(3, 10)

	// Definir procesador
	processor := func(task Task) Result {
		time.Sleep(200 * time.Millisecond)
		return Result{
			Output: fmt.Sprintf("Processed: %s", task.Data),
			Error:  nil,
		}
	}

	// Iniciar workers
	pool.Start(processor)

	// Enviar tareas
	for i := 1; i <= 5; i++ {
		pool.Submit(Task{
			ID:   i,
			Data: fmt.Sprintf("Task %d", i),
		})
	}

	// Cerrar y recibir resultados
	pool.Close()
	for result := range pool.Results() {
		fmt.Printf("Result: %s\n", result.Output)
	}
	fmt.Println()
}

// ============================================================================
// EJEMPLO PRÁCTICO: PROCESAR MILES DE TAREAS
// ============================================================================
// Este es el ejemplo sugerido en el temario

func processThousandsOfTasks() {
	fmt.Println("=== Process Thousands of Tasks ===")

	numWorkers := 10
	numTasks := 1000

	jobs := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// Workers
	var wg sync.WaitGroup
	startTime := time.Now()

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for taskID := range jobs {
				// Simular trabajo (procesar número)
				result := taskID * 2
				time.Sleep(10 * time.Millisecond) // Simular procesamiento
				results <- result
			}
		}()
	}

	// Enviar todas las tareas
	for i := 1; i <= numTasks; i++ {
		jobs <- i
	}
	close(jobs)

	// Cerrar results cuando terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recibir resultados
	count := 0
	for range results {
		count++
		if count%100 == 0 {
			fmt.Printf("Processed %d tasks...\n", count)
		}
	}

	elapsed := time.Since(startTime)
	fmt.Printf("Processed %d tasks in %v\n", count, elapsed)
	fmt.Printf("Throughput: %.2f tasks/second\n", float64(count)/elapsed.Seconds())
	fmt.Println()
}

// ============================================================================
// WORKER POOL CON PRIORIDADES
// ============================================================================

type PriorityTask struct {
	Task
	Priority int // Mayor número = mayor prioridad
}

func priorityWorkerPool() {
	fmt.Println("=== Priority Worker Pool ===")

	// Usar buffered channel para prioridades
	highPriority := make(chan PriorityTask, 10)
	lowPriority := make(chan PriorityTask, 10)

	// Worker que procesa por prioridad
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case task := <-highPriority:
				fmt.Printf("HIGH PRIORITY: Processing %s\n", task.Data)
				time.Sleep(100 * time.Millisecond)
			case task := <-lowPriority:
				fmt.Printf("LOW PRIORITY: Processing %s\n", task.Data)
				time.Sleep(100 * time.Millisecond)
			default:
				time.Sleep(50 * time.Millisecond)
				// Verificar si ambos están vacíos y cerrar
				if len(highPriority) == 0 && len(lowPriority) == 0 {
					return
				}
			}
		}
	}()

	// Enviar tareas con diferentes prioridades
	lowPriority <- PriorityTask{Task: Task{ID: 1, Data: "Low 1"}, Priority: 1}
	highPriority <- PriorityTask{Task: Task{ID: 2, Data: "High 1"}, Priority: 10}
	lowPriority <- PriorityTask{Task: Task{ID: 3, Data: "Low 2"}, Priority: 1}
	highPriority <- PriorityTask{Task: Task{ID: 4, Data: "High 2"}, Priority: 10}

	time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println()
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	fmt.Println("=== CONCURRENCIA: WORKER POOLS ===\n")

	basicWorkerPool()
	workerPoolWithContext()
	reusableWorkerPool()
	processThousandsOfTasks()
	priorityWorkerPool()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

