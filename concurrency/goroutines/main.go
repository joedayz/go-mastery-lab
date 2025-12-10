package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================================
// GOROUTINES
// ============================================================================
// Las goroutines son funciones que se ejecutan concurrentemente
// Son MUCHO más ligeras que los threads en Java
// - Threads en Java: ~1-2MB de stack cada uno
// - Goroutines en Go: ~2KB inicialmente, crece según necesidad
// Puedes tener millones de goroutines ejecutándose simultáneamente
// ============================================================================

// Función simple que se ejecutará como goroutine
func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello from %s (iteration %d)\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

// ============================================================================
// GOROUTINES BÁSICAS
// ============================================================================

func basicGoroutines() {
	fmt.Println("=== Goroutines Básicas ===")

	// Ejecutar función como goroutine (no bloquea)
	go sayHello("Goroutine 1")
	go sayHello("Goroutine 2")

	// Esperar un poco para que las goroutines terminen
	// En producción, usarías channels o sync.WaitGroup
	time.Sleep(1 * time.Second)
	fmt.Println()
}

// ============================================================================
// GOROUTINES CON ANÓNIMAS
// ============================================================================

func anonymousGoroutines() {
	fmt.Println("=== Goroutines Anónimas ===")

	// Función anónima como goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("Hello from anonymous goroutine")

	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

// ============================================================================
// GOROUTINES CON CLOSURES
// ============================================================================
// CUIDADO: Las goroutines capturan variables por referencia
// Similar a lambdas en Java, pero más cuidadoso con el scope

func goroutinesWithClosures() {
	fmt.Println("=== Goroutines con Closures ===")

	// PROBLEMA COMÚN: Loop variable capturada incorrectamente
	fmt.Println("Problema común (incorrecto):")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("Incorrect: i = %d\n", i) // i puede ser cualquier valor
		}()
	}
	time.Sleep(100 * time.Millisecond)

	// SOLUCIÓN 1: Pasar como parámetro
	fmt.Println("\nSolución 1 (pasar como parámetro):")
	for i := 0; i < 3; i++ {
		go func(val int) {
			fmt.Printf("Correct: val = %d\n", val)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)

	// SOLUCIÓN 2: Crear copia local
	fmt.Println("\nSolución 2 (copia local):")
	for i := 0; i < 3; i++ {
		i := i // Crear nueva variable en cada iteración
		go func() {
			fmt.Printf("Correct: i = %d\n", i)
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

// ============================================================================
// SINCRONIZACIÓN CON SYNC.WAITGROUP
// ============================================================================
// WaitGroup es la forma idiomática de esperar múltiples goroutines
// Similar a CountDownLatch en Java

func waitGroupExample() {
	fmt.Println("=== WaitGroup ===")

	var wg sync.WaitGroup

	// Lanzar múltiples goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1) // Incrementar contador
		go func(id int) {
			defer wg.Done() // Decrementar contador cuando termine
			fmt.Printf("Goroutine %d: Starting\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Goroutine %d: Finished\n", id)
		}(i)
	}

	// Esperar a que todas terminen
	fmt.Println("Waiting for all goroutines to finish...")
	wg.Wait()
	fmt.Println("All goroutines finished!")
	fmt.Println()
}

// ============================================================================
// EJEMPLO PRÁCTICO: PROCESAR TAREAS EN PARALELO
// ============================================================================

type Task struct {
	ID   int
	Data string
}

func processTask(task Task) {
	fmt.Printf("Processing task %d: %s\n", task.ID, task.Data)
	time.Sleep(500 * time.Millisecond) // Simular trabajo
	fmt.Printf("Task %d completed\n", task.ID)
}

func processTasksConcurrently() {
	fmt.Println("=== Procesar Tareas Concurrentemente ===")

	tasks := []Task{
		{ID: 1, Data: "Task 1"},
		{ID: 2, Data: "Task 2"},
		{ID: 3, Data: "Task 3"},
		{ID: 4, Data: "Task 4"},
		{ID: 5, Data: "Task 5"},
	}

	var wg sync.WaitGroup

	start := time.Now()

	// Procesar todas las tareas en paralelo
	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			processTask(t)
		}(task)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("All tasks completed in %v\n", elapsed)
	fmt.Println()
}

// ============================================================================
// EJEMPLO: SIMULACIÓN DE SERVIDOR WEB
// ============================================================================

type Request struct {
	ID   int
	Path string
}

func handleRequest(req Request) {
	fmt.Printf("[Request %d] Handling %s\n", req.ID, req.Path)
	time.Sleep(200 * time.Millisecond) // Simular procesamiento
	fmt.Printf("[Request %d] Completed\n", req.ID)
}

func simulateWebServer() {
	fmt.Println("=== Simulación de Servidor Web ===")

	requests := []Request{
		{ID: 1, Path: "/api/users"},
		{ID: 2, Path: "/api/products"},
		{ID: 3, Path: "/api/orders"},
		{ID: 4, Path: "/api/payments"},
		{ID: 5, Path: "/api/reports"},
	}

	var wg sync.WaitGroup

	// Simular servidor que maneja requests concurrentemente
	for _, req := range requests {
		wg.Add(1)
		go func(r Request) {
			defer wg.Done()
			handleRequest(r)
		}(req)
	}

	wg.Wait()
	fmt.Println("All requests handled")
	fmt.Println()
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	fmt.Println("=== CONCURRENCIA: GOROUTINES ===\n")

	basicGoroutines()
	anonymousGoroutines()
	goroutinesWithClosures()
	waitGroupExample()
	processTasksConcurrently()
	simulateWebServer()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

