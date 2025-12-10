package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================================
// SYNC - PRIMITIVAS DE SINCRONIZACIÓN
// ============================================================================
// sync proporciona primitivas de bajo nivel para sincronización
// Similar a java.util.concurrent en Java
// ============================================================================

// ============================================================================
// SYNC.MUTEX - EXCLUSIÓN MUTUA
// ============================================================================
// Mutex protege secciones críticas
// Similar a synchronized en Java o ReentrantLock

func mutexExample() {
	fmt.Println("=== Mutex Example ===")

	var mu sync.Mutex
	counter := 0

	// Incrementar contador de forma segura
	increment := func() {
		mu.Lock()
		defer mu.Unlock() // Siempre usar defer para unlock
		counter++
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter: %d (should be 1000)\n", counter)
	fmt.Println()
}

// ============================================================================
// SYNC.RWMUTEX - READ-WRITE MUTEX
// ============================================================================
// Permite múltiples lectores simultáneos o un escritor exclusivo
// Similar a ReadWriteLock en Java

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()         // Lock para lectura (múltiples lectores permitidos)
	defer sm.mu.RUnlock() // Unlock para lectura
	value, exists := sm.data[key]
	return value, exists
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()         // Lock para escritura (exclusivo)
	defer sm.mu.Unlock() // Unlock para escritura
	sm.data[key] = value
}

func rwMutexExample() {
	fmt.Println("=== RWMutex Example ===")

	sm := NewSafeMap()

	// Escritores
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sm.Set(fmt.Sprintf("key%d", id), id)
		}(i)
	}

	// Lectores (pueden leer simultáneamente)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			value, exists := sm.Get(fmt.Sprintf("key%d", id))
			if exists {
				fmt.Printf("Read key%d: %d\n", id, value)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println()
}

// ============================================================================
// SYNC.WAITGROUP
// ============================================================================
// Esperar a que múltiples goroutines terminen
// Similar a CountDownLatch en Java

func waitGroupExample() {
	fmt.Println("=== WaitGroup Example ===")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d: Starting\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Worker %d: Finished\n", id)
		}(i)
	}

	fmt.Println("Waiting for all workers...")
	wg.Wait()
	fmt.Println("All workers finished!")
	fmt.Println()
}

// ============================================================================
// SYNC.ONCE - EJECUTAR UNA VEZ
// ============================================================================
// Garantiza que una función se ejecute solo una vez
// Similar a lazy initialization en Java

func onceExample() {
	fmt.Println("=== Once Example ===")

	var once sync.Once
	initialize := func() {
		fmt.Println("Initializing (should only see this once)")
	}

	// Intentar ejecutar múltiples veces
	for i := 0; i < 5; i++ {
		once.Do(initialize)
	}
	fmt.Println()
}

// ============================================================================
// SYNC.COND - CONDITION VARIABLES
// ============================================================================
// Permite que goroutines esperen por condiciones
// Similar a Condition en Java

func condExample() {
	fmt.Println("=== Cond Example ===")

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	// Worker que espera condición
	go func() {
		mu.Lock()
		for !ready {
			fmt.Println("Worker: Waiting for condition...")
			cond.Wait() // Libera el lock y espera
		}
		fmt.Println("Worker: Condition met, proceeding!")
		mu.Unlock()
	}()

	// Simular trabajo antes de señalizar
	time.Sleep(1 * time.Second)

	// Señalizar condición
	mu.Lock()
	ready = true
	fmt.Println("Main: Signaling condition...")
	cond.Signal() // Despierta una goroutine esperando
	mu.Unlock()

	time.Sleep(500 * time.Millisecond)
	fmt.Println()
}

// ============================================================================
// SYNC.POOL - OBJECT POOLING
// ============================================================================
// Pool de objetos reutilizables para reducir allocations
// Similar a ObjectPool pattern en Java

func poolExample() {
	fmt.Println("=== Pool Example ===")

	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new buffer")
			return make([]byte, 1024)
		},
	}

	// Obtener del pool
	buf1 := pool.Get().([]byte)
	fmt.Printf("Got buffer from pool: len=%d\n", len(buf1))

	// Devolver al pool
	pool.Put(buf1)

	// Obtener de nuevo (puede ser el mismo objeto)
	buf2 := pool.Get().([]byte)
	fmt.Printf("Got buffer from pool again: len=%d\n", len(buf2))

	pool.Put(buf2)
	fmt.Println()
}

// ============================================================================
// ATOMIC OPERATIONS
// ============================================================================
// Operaciones atómicas sin locks
// Similar a AtomicInteger, AtomicLong en Java

func atomicExample() {
	fmt.Println("=== Atomic Operations ===")

	var counter int64

	// Incrementar atómicamente
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("Counter: %d (should be 1000)\n", atomic.LoadInt64(&counter))

	// Compare and swap
	oldValue := atomic.LoadInt64(&counter)
	newValue := oldValue + 100
	if atomic.CompareAndSwapInt64(&counter, oldValue, newValue) {
		fmt.Printf("CAS succeeded: %d -> %d\n", oldValue, newValue)
	}
	fmt.Println()
}

// ============================================================================
// EJEMPLO PRÁCTICO: CACHE THREAD-SAFE
// ============================================================================

type Cache struct {
	mu    sync.RWMutex
	data  map[string]interface{}
	stats struct {
		hits   int64
		misses int64
	}
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	value, exists := c.data[key]
	c.mu.RUnlock()

	if exists {
		atomic.AddInt64(&c.stats.hits, 1)
	} else {
		atomic.AddInt64(&c.stats.misses, 1)
	}

	return value, exists
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Stats() (hits, misses int64) {
	return atomic.LoadInt64(&c.stats.hits), atomic.LoadInt64(&c.stats.misses)
}

func cacheExample() {
	fmt.Println("=== Cache Example ===")

	cache := NewCache()

	// Escritores
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cache.Set(fmt.Sprintf("key%d", id), fmt.Sprintf("value%d", id))
		}(i)
	}

	// Lectores
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id%10)
			value, exists := cache.Get(key)
			if exists {
				fmt.Printf("Cache hit: %s = %s\n", key, value)
			} else {
				fmt.Printf("Cache miss: %s\n", key)
			}
		}(i)
	}

	wg.Wait()
	hits, misses := cache.Stats()
	fmt.Printf("Cache stats: hits=%d, misses=%d\n", hits, misses)
	fmt.Println()
}

// ============================================================================
// DETECTAR DATA RACES
// ============================================================================
// Go tiene un race detector incorporado
// Ejecutar con: go run -race main.go
// O: go test -race

func dataRaceExample() {
	fmt.Println("=== Data Race Example (BAD CODE) ===")
	fmt.Println("This code has a data race!")
	fmt.Println("Run with: go run -race main.go to detect it")
	fmt.Println()

	var counter int

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// DATA RACE: Múltiples goroutines escriben sin lock
			counter++
		}()
	}

	wg.Wait()
	fmt.Printf("Counter: %d (may not be 10 due to race condition)\n", counter)
	fmt.Println()
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	fmt.Println("=== CONCURRENCIA: SYNC ===\n")

	mutexExample()
	rwMutexExample()
	waitGroupExample()
	onceExample()
	condExample()
	poolExample()
	atomicExample()
	cacheExample()
	dataRaceExample()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

