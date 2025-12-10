package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Importar pprof automáticamente
	"runtime"
	"time"
)

// ============================================================================
// PROFILING CON PPROF
// ============================================================================
// pprof es la herramienta de profiling de Go
// Similar a JProfiler o VisualVM en Java
// ============================================================================

func main() {
	// Iniciar servidor HTTP para pprof
	// Acceder a: http://localhost:6060/debug/pprof/
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("pprof server started at http://localhost:6060/debug/pprof/")
	fmt.Println("Available profiles:")
	fmt.Println("  - CPU:   go tool pprof http://localhost:6060/debug/pprof/profile")
	fmt.Println("  - Heap:  go tool pprof http://localhost:6060/debug/pprof/heap")
	fmt.Println("  - Mutex: go tool pprof http://localhost:6060/debug/pprof/mutex")

	// Simular carga de trabajo
	for {
		cpuIntensiveTask()
		memoryIntensiveTask()
		time.Sleep(1 * time.Second)
	}
}

func cpuIntensiveTask() {
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
}

func memoryIntensiveTask() {
	// Crear slice grande
	data := make([]byte, 1024*1024) // 1MB
	for i := range data {
		data[i] = byte(i % 256)
	}
	// Forzar GC periódicamente
	if time.Now().Unix()%10 == 0 {
		runtime.GC()
	}
}

