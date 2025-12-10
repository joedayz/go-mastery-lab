package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
	"time"
)

// ============================================================================
// PAQUETES IMPORTANTES DE GO
// ============================================================================
// Go viene con una biblioteca estándar muy completa
// Estos son algunos de los paquetes más importantes para nivel Senior
// ============================================================================

// ============================================================================
// CONTEXT - Manejo de cancelación y timeouts
// ============================================================================
// context es fundamental para manejar cancelación, timeouts y valores
// en requests HTTP, operaciones de base de datos, etc.
// Similar a CompletableFuture en Java, pero más integrado

func demonstrateContext() {
	// Context con timeout
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

	// Context con cancelación manual
	ctx2, cancel2 := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		cancel2() // Cancelar después de 1 segundo
	}()

	<-ctx2.Done()
	fmt.Printf("Context cancelled: %v\n", ctx2.Err())

	// Context con valores (útil para pasar datos entre funciones)
	ctx3 := context.WithValue(context.Background(), "userID", 123)
	ctx3 = context.WithValue(ctx3, "requestID", "req-456")

	userID := ctx3.Value("userID").(int)
	requestID := ctx3.Value("requestID").(string)
	fmt.Printf("User ID: %d, Request ID: %s\n", userID, requestID)
}

// ============================================================================
// IO - Operaciones de entrada/salida
// ============================================================================
// io proporciona interfaces y funciones para I/O
// Similar a java.io en Java

func demonstrateIO() {
	// io.Reader - interfaz para leer datos
	file, err := os.Open("fundamentals/packages/main.go")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Leer todo el contenido
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("Read %d bytes from file\n", len(data))

	// io.Writer - interfaz para escribir datos
	output := os.Stdout
	io.WriteString(output, "Hello from io.WriteString!\n")

	// io.Copy - copiar de Reader a Writer
	reader := io.NopCloser(io.Reader(file))
	written, err := io.Copy(os.Stdout, reader)
	if err != nil {
		fmt.Printf("Error copying: %v\n", err)
	} else {
		fmt.Printf("Copied %d bytes\n", written)
	}
}

// ============================================================================
// FMT - Formateo e impresión
// ============================================================================
// fmt es similar a System.out.println y String.format en Java

func demonstrateFMT() {
	name := "John"
	age := 30
	balance := 1234.56

	// Printf - formateo con formato
	fmt.Printf("Name: %s, Age: %d, Balance: $%.2f\n", name, age, balance)

	// Sprintf - formateo a string
	message := fmt.Sprintf("User %s is %d years old", name, age)
	fmt.Println(message)

	// Scanf - leer entrada formateada (similar a Scanner en Java)
	fmt.Println("Enter name and age:")
	var inputName string
	var inputAge int
	fmt.Scanf("%s %d", &inputName, &inputAge)
	fmt.Printf("You entered: %s, %d\n", inputName, inputAge)

	// Fprintf - escribir a un Writer
	fmt.Fprintf(os.Stderr, "This is an error message\n")
}

// ============================================================================
// SYNC - Sincronización y concurrencia
// ============================================================================
// sync proporciona primitivas de sincronización
// Similar a java.util.concurrent en Java

func demonstrateSync() {
	// sync.Mutex - mutex para exclusión mutua
	var mu sync.Mutex
	counter := 0

	// Incrementar contador de forma segura
	increment := func() {
		mu.Lock()
		defer mu.Unlock()
		counter++
	}

	// Ejecutar múltiples goroutines
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter after 10 increments: %d\n", counter)

	// sync.RWMutex - mutex de lectura/escritura
	var rwmu sync.RWMutex
	data := make(map[string]int)

	// Escritura (exclusiva)
	write := func(key string, value int) {
		rwmu.Lock()
		defer rwmu.Unlock()
		data[key] = value
		fmt.Printf("Wrote %s = %d\n", key, value)
	}

	// Lectura (múltiples lectores simultáneos)
	read := func(key string) int {
		rwmu.RLock()
		defer rwmu.RUnlock()
		return data[key]
	}

	write("a", 1)
	write("b", 2)
	fmt.Printf("Read a: %d\n", read("a"))

	// sync.WaitGroup - esperar a que múltiples goroutines terminen
	var wg2 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}
	wg2.Wait()
	fmt.Println("All goroutines completed")

	// sync.Once - ejecutar algo solo una vez
	var once sync.Once
	initialize := func() {
		fmt.Println("Initializing (should only see this once)")
	}

	for i := 0; i < 5; i++ {
		once.Do(initialize)
	}
}

// ============================================================================
// REFLECT - Reflexión en tiempo de ejecución
// ============================================================================
// reflect permite inspeccionar tipos en tiempo de ejecución
// Similar a java.lang.reflect en Java, pero más limitado

func demonstrateReflect() {
	// Obtener tipo de un valor
	var x int = 42
	t := reflect.TypeOf(x)
	fmt.Printf("Type of x: %s\n", t)

	// Obtener valor
	v := reflect.ValueOf(x)
	fmt.Printf("Value of x: %v\n", v.Int())

	// Inspeccionar struct
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}
	pType := reflect.TypeOf(p)
	pValue := reflect.ValueOf(p)

	fmt.Printf("Struct type: %s\n", pType)
	fmt.Printf("Number of fields: %d\n", pType.NumField())

	for i := 0; i < pType.NumField(); i++ {
		field := pType.Field(i)
		value := pValue.Field(i)
		fmt.Printf("  Field %d: %s (%s) = %v\n", i, field.Name, field.Type, value.Interface())
	}

	// Modificar valores (debe ser un pointer)
	p2 := &Person{Name: "Jane", Age: 25}
	p2Value := reflect.ValueOf(p2).Elem()
	ageField := p2Value.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(26)
		fmt.Printf("Modified age: %d\n", p2.Age)
	}
}

// ============================================================================
// EJEMPLO PRÁCTICO: COMBINANDO MÚLTIPLES PAQUETES
// ============================================================================

type DataProcessor struct {
	mu    sync.RWMutex
	data  map[string]interface{}
	ctx   context.Context
	cancel context.CancelFunc
}

func NewDataProcessor() *DataProcessor {
	ctx, cancel := context.WithCancel(context.Background())
	return &DataProcessor{
		data:   make(map[string]interface{}),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (dp *DataProcessor) Set(key string, value interface{}) {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	dp.data[key] = value
}

func (dp *DataProcessor) Get(key string) (interface{}, bool) {
	dp.mu.RLock()
	defer dp.mu.RUnlock()
	value, exists := dp.data[key]
	return value, exists
}

func (dp *DataProcessor) ProcessWithTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(dp.ctx, timeout)
	defer cancel()

	done := make(chan error)
	go func() {
		// Simular procesamiento
		time.Sleep(2 * time.Second)
		done <- nil
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (dp *DataProcessor) Stop() {
	dp.cancel()
}

// ============================================================================
// MAIN - Ejemplos de uso
// ============================================================================

func main() {
	fmt.Println("=== FUNDAMENTOS: PAQUETES IMPORTANTES ===\n")

	// 1. Context
	fmt.Println("1. Context (cancelación y timeouts):")
	demonstrateContext()
	fmt.Println()

	// 2. IO
	fmt.Println("2. IO (entrada/salida):")
	demonstrateIO()
	fmt.Println()

	// 3. FMT
	fmt.Println("3. FMT (formateo):")
	demonstrateFMT()
	fmt.Println()

	// 4. Sync
	fmt.Println("4. Sync (sincronización):")
	demonstrateSync()
	fmt.Println()

	// 5. Reflect
	fmt.Println("5. Reflect (reflexión):")
	demonstrateReflect()
	fmt.Println()

	// 6. Ejemplo práctico combinado
	fmt.Println("6. Ejemplo práctico combinado:")
	processor := NewDataProcessor()
	processor.Set("key1", "value1")
	processor.Set("key2", 42)

	if value, exists := processor.Get("key1"); exists {
		fmt.Printf("Retrieved: %v\n", value)
	}

	if err := processor.ProcessWithTimeout(1 * time.Second); err != nil {
		fmt.Printf("Processing error (expected timeout): %v\n", err)
	}

	processor.Stop()
	fmt.Println()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

