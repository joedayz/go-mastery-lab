package main

import (
	"fmt"
)

// ============================================================================
// MÉTODOS Y RECEPTORES (VALUE VS POINTER RECEIVER)
// ============================================================================
// En Go, puedes definir métodos en cualquier tipo (no solo structs)
// Los métodos pueden tener value receivers o pointer receivers
// Esta es una decisión importante que afecta el comportamiento
// ============================================================================

// ============================================================================
// VALUE RECEIVER
// ============================================================================
// Recibe una COPIA del valor
// Útil cuando:
// - El struct es pequeño
// - No necesitas modificar el struct original
// - Quieres inmutabilidad

type Counter struct {
	value int
}

// Increment con value receiver - NO modifica el original
func (c Counter) Increment() {
	c.value++ // Esto solo modifica la copia
	fmt.Printf("Inside Increment (value receiver): %d\n", c.value)
}

// GetValue con value receiver - solo lee, no modifica
func (c Counter) GetValue() int {
	return c.value
}

// ============================================================================
// POINTER RECEIVER
// ============================================================================
// Recibe una REFERENCIA al valor
// Útil cuando:
// - El struct es grande (evita copias costosas)
// - Necesitas modificar el struct original
// - Quieres consistencia (si un método modifica, todos deberían usar pointer)

type CounterPtr struct {
	value int
}

// Increment con pointer receiver - SÍ modifica el original
func (c *CounterPtr) Increment() {
	c.value++ // Esto modifica el original
	fmt.Printf("Inside Increment (pointer receiver): %d\n", c.value)
}

// GetValue con pointer receiver
func (c *CounterPtr) GetValue() int {
	return c.value
}

// ============================================================================
// CUÁNDO USAR VALUE VS POINTER RECEIVER
// ============================================================================
// Regla general:
// - Si el método modifica el receiver → usa pointer receiver
// - Si el struct es grande → usa pointer receiver
// - Si necesitas consistencia → usa pointer receiver para todos los métodos
// - Si solo lees y el struct es pequeño → value receiver está bien

// ============================================================================
// EJEMPLO: STRUCT PEQUEÑO CON VALUE RECEIVER
// ============================================================================

type Point struct {
	X, Y float64
}

// Distance calcula la distancia desde el origen (no modifica Point)
func (p Point) Distance() float64 {
	return p.X*p.X + p.Y*p.Y
}

// ============================================================================
// EJEMPLO: STRUCT GRANDE CON POINTER RECEIVER
// ============================================================================

type LargeStruct struct {
	data [1000]int
	name string
}

// UpdateName modifica el struct, usa pointer receiver
func (ls *LargeStruct) UpdateName(name string) {
	ls.name = name
}

// GetName solo lee, pero usa pointer receiver por consistencia
func (ls *LargeStruct) GetName() string {
	return ls.name
}

// ============================================================================
// MÉTODOS EN TIPOS NO-STRUCT
// ============================================================================
// Puedes definir métodos en cualquier tipo definido por ti

// MyInt es un tipo personalizado basado en int
type MyInt int

// Double duplica el valor (value receiver porque int es pequeño)
func (m MyInt) Double() MyInt {
	return m * 2
}

// String implementa Stringer para MyInt
func (m MyInt) String() string {
	return fmt.Sprintf("MyInt(%d)", int(m))
}

// ============================================================================
// MÉTODOS EN SLICES (TIPO PERSONALIZADO)
// ============================================================================

type IntSlice []int

// Sum calcula la suma de todos los elementos
func (is IntSlice) Sum() int {
	sum := 0
	for _, v := range is {
		sum += v
	}
	return sum
}

// Append agrega un valor (retorna nuevo slice, no modifica el original)
func (is IntSlice) Append(value int) IntSlice {
	return append(is, value)
}

// ============================================================================
// EJEMPLO PRÁCTICO: BANCO CON CUENTAS
// ============================================================================

type Account struct {
	ID      string
	Balance float64
	Owner   string
}

// Deposit deposita dinero (modifica, usa pointer receiver)
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	a.Balance += amount
	return nil
}

// Withdraw retira dinero (modifica, usa pointer receiver)
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if a.Balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

// GetBalance obtiene el balance (lee, pero usa pointer por consistencia)
func (a *Account) GetBalance() float64 {
	return a.Balance
}

// String implementa Stringer (value receiver porque solo lee)
func (a Account) String() string {
	return fmt.Sprintf("Account{ID: %s, Owner: %s, Balance: $%.2f}", a.ID, a.Owner, a.Balance)
}

// ============================================================================
// MÉTODOS CON MÚLTIPLES RECEIVERS (NO PERMITIDO)
// ============================================================================
// No puedes tener el mismo método con value y pointer receiver
// Debes elegir uno y ser consistente

// ============================================================================
// GO AUTOMÁTICAMENTE CONVIERTE ENTRE VALUE Y POINTER
// ============================================================================
// Si tienes un método con pointer receiver (*T), Go automáticamente
// toma la dirección si pasas un value
// Si tienes un método con value receiver (T), Go automáticamente
// dereferencia si pasas un pointer

func demonstrateAutomaticConversion() {
	// Crear un Account
	acc := Account{ID: "ACC001", Balance: 1000.0, Owner: "John"}

	// Llamar método con pointer receiver usando un value
	// Go automáticamente convierte: (&acc).Deposit(100)
	acc.Deposit(100)
	fmt.Printf("After deposit: %s\n", acc)

	// Crear un pointer a Account
	accPtr := &Account{ID: "ACC002", Balance: 500.0, Owner: "Jane"}

	// Llamar método con value receiver usando un pointer
	// Go automáticamente convierte: (*accPtr).String()
	fmt.Printf("Account: %s\n", accPtr.String())
}

// ============================================================================
// MAIN - Ejemplos de uso
// ============================================================================

func main() {
	fmt.Println("=== FUNDAMENTOS: MÉTODOS Y RECEPTORES ===\n")

	// 1. Value receiver vs Pointer receiver
	fmt.Println("1. Value Receiver (no modifica el original):")
	counter := Counter{value: 0}
	fmt.Printf("Before: %d\n", counter.GetValue())
	counter.Increment()
	fmt.Printf("After: %d (no cambió)\n", counter.GetValue())
	fmt.Println()

	fmt.Println("2. Pointer Receiver (sí modifica el original):")
	counterPtr := CounterPtr{value: 0}
	fmt.Printf("Before: %d\n", counterPtr.GetValue())
	counterPtr.Increment()
	fmt.Printf("After: %d (sí cambió)\n", counterPtr.GetValue())
	fmt.Println()

	// 2. Métodos en tipos no-struct
	fmt.Println("3. Métodos en tipos personalizados:")
	myInt := MyInt(5)
	fmt.Printf("Original: %s\n", myInt)
	fmt.Printf("Doubled: %s\n", myInt.Double())
	fmt.Println()

	// 3. Métodos en slices personalizados
	fmt.Println("4. Métodos en slices personalizados:")
	numbers := IntSlice{1, 2, 3, 4, 5}
	fmt.Printf("Sum of %v: %d\n", numbers, numbers.Sum())
	numbers = numbers.Append(6)
	fmt.Printf("After append: %v\n", numbers)
	fmt.Println()

	// 4. Ejemplo práctico: Cuenta bancaria
	fmt.Println("5. Ejemplo práctico: Cuenta bancaria:")
	account := &Account{
		ID:      "ACC001",
		Balance: 1000.0,
		Owner:   "John Doe",
	}
	fmt.Printf("Initial: %s\n", account)

	if err := account.Deposit(500.0); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After deposit: %s\n", account)
	}

	if err := account.Withdraw(200.0); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After withdrawal: %s\n", account)
	}

	if err := account.Withdraw(2000.0); err != nil {
		fmt.Printf("Error (expected): %v\n", err)
	}
	fmt.Println()

	// 5. Conversión automática
	fmt.Println("6. Conversión automática entre value y pointer:")
	demonstrateAutomaticConversion()
	fmt.Println()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

