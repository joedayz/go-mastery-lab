package main

import (
	"fmt"
	"time"
)

// ============================================================================
// CHANNELS (BUFFERED VS UNBUFFERED)
// ============================================================================
// Channels son la forma idiomática de comunicar goroutines en Go
// Similar a BlockingQueue en Java, pero integrado en el lenguaje
// 
// Unbuffered: bloquea hasta que hay un receptor (sincronización)
// Buffered: tiene capacidad, solo bloquea cuando está lleno
// ============================================================================

// ============================================================================
// CHANNELS UNBUFFERED
// ============================================================================
// Unbuffered channels bloquean hasta que hay un receptor
// Son perfectos para sincronización

func unbufferedChannels() {
	fmt.Println("=== Unbuffered Channels ===")

	// Crear channel unbuffered
	ch := make(chan string)

	// Goroutine que envía
	go func() {
		fmt.Println("Sending message...")
		ch <- "Hello from goroutine"
		fmt.Println("Message sent!")
	}()

	// Esperar un poco (para demostrar el bloqueo)
	time.Sleep(500 * time.Millisecond)

	// Recibir mensaje (desbloquea el sender)
	msg := <-ch
	fmt.Printf("Received: %s\n", msg)
	fmt.Println()
}

// ============================================================================
// CHANNELS BUFFERED
// ============================================================================
// Buffered channels tienen capacidad
// Solo bloquean cuando están llenos (sender) o vacíos (receiver)

func bufferedChannels() {
	fmt.Println("=== Buffered Channels ===")

	// Crear channel con buffer de tamaño 3
	ch := make(chan int, 3)

	// Enviar múltiples valores sin bloqueo (hasta que el buffer esté lleno)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Sent 3 values to buffered channel")

	// El siguiente envío bloquearía porque el buffer está lleno
	// ch <- 4 // Esto bloquearía

	// Recibir valores
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Println()
}

// ============================================================================
// CHANNELS CON RANGE
// ============================================================================
// Puedes iterar sobre un channel hasta que se cierre

func channelRange() {
	fmt.Println("=== Channel Range ===")

	ch := make(chan int)

	// Goroutine que envía valores y cierra el channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch) // Cerrar channel cuando terminamos
	}()

	// Iterar sobre el channel
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println("Channel closed")
	fmt.Println()
}

// ============================================================================
// CHANNELS DIRECCIONALES
// ============================================================================
// Puedes especificar si un channel solo envía o solo recibe

func directionalChannels() {
	fmt.Println("=== Directional Channels ===")

	ch := make(chan string)

	// Función que solo envía (chan<-)
	sender := func(ch chan<- string) {
		ch <- "Hello"
		ch <- "World"
		close(ch)
	}

	// Función que solo recibe (<-chan)
	receiver := func(ch <-chan string) {
		for msg := range ch {
			fmt.Printf("Received: %s\n", msg)
		}
	}

	go sender(ch)
	receiver(ch)
	fmt.Println()
}

// ============================================================================
// SELECT STATEMENT
// ============================================================================
// Select permite esperar en múltiples channels simultáneamente
// Similar a Selector en Java NIO, pero más simple

func selectStatement() {
	fmt.Println("=== Select Statement ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine que envía a ch1 después de 1 segundo
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	// Goroutine que envía a ch2 después de 2 segundos
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	// Select espera en ambos channels
	// Ejecuta el primer case que esté listo
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received from ch1: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received from ch2: %s\n", msg2)
	}

	fmt.Println()
}

// ============================================================================
// SELECT CON DEFAULT (NON-BLOCKING)
// ============================================================================
// Default case hace que select no bloquee

func selectWithDefault() {
	fmt.Println("=== Select with Default (Non-blocking) ===")

	ch := make(chan string)

	// Intentar recibir sin bloquear
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message available (non-blocking)")
	}

	// Enviar sin bloquear (solo funciona con buffered channels)
	bufferedCh := make(chan string, 1)
	select {
	case bufferedCh <- "Hello":
		fmt.Println("Sent to buffered channel")
	default:
		fmt.Println("Channel full, couldn't send")
	}

	fmt.Println()
}

// ============================================================================
// SELECT CON TIMEOUT
// ============================================================================
// Usar time.After para timeouts

func selectWithTimeout() {
	fmt.Println("=== Select with Timeout ===")

	ch := make(chan string)

	// Goroutine que tarda mucho
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Slow message"
	}()

	// Esperar con timeout de 1 segundo
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! No message received")
	}

	fmt.Println()
}

// ============================================================================
// EJEMPLO PRÁCTICO: WORKER POOL SIMPLE
// ============================================================================

func workerPoolExample() {
	fmt.Println("=== Worker Pool Example ===")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Crear 3 workers
	for w := 1; w <= 3; w++ {
		go func(id int) {
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(500 * time.Millisecond) // Simular trabajo
				results <- job * 2
			}
		}(w)
	}

	// Enviar trabajos
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Recibir resultados
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Println()
}

// ============================================================================
// EJEMPLO: FAN-OUT / FAN-IN PATTERN
// ============================================================================
// Fan-out: múltiples workers procesan de un channel
// Fan-in: combinar resultados de múltiples channels en uno

func fanOutFanIn() {
	fmt.Println("=== Fan-Out / Fan-In Pattern ===")

	// Channel de entrada
	input := make(chan int)

	// Fan-out: múltiples workers procesan
	worker1 := make(chan int)
	worker2 := make(chan int)

	go func() {
		for val := range input {
			worker1 <- val * 2
		}
		close(worker1)
	}()

	go func() {
		for val := range input {
			worker2 <- val * 3
		}
		close(worker2)
	}()

	// Fan-in: combinar resultados
	output := make(chan int)
	go func() {
		for {
			select {
			case val, ok := <-worker1:
				if !ok {
					worker1 = nil
				} else {
					output <- val
				}
			case val, ok := <-worker2:
				if !ok {
					worker2 = nil
				} else {
					output <- val
				}
			}
			if worker1 == nil && worker2 == nil {
				close(output)
				return
			}
		}
	}()

	// Enviar datos
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	// Recibir resultados combinados
	for result := range output {
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Println()
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	fmt.Println("=== CONCURRENCIA: CHANNELS ===\n")

	unbufferedChannels()
	bufferedChannels()
	channelRange()
	directionalChannels()
	selectStatement()
	selectWithDefault()
	selectWithTimeout()
	workerPoolExample()
	fanOutFanIn()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

