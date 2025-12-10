package main

import (
	"fmt"
	"time"
)

// ============================================================================
// TIPOS Y MANEJO DE STRUCTS
// ============================================================================
// En Go, los structs son similares a las clases en Java, pero más simples.
// No tienen herencia, pero puedes usar composición (embedding).
// ============================================================================

// Person representa una persona básica
// En Java sería: public class Person { ... }
type Person struct {
	// Campos públicos (empiezan con mayúscula) - similar a public en Java
	Name string
	Age  int

	// Campos privados (empiezan con minúscula) - similar a private en Java
	email string

	// Puedes agregar tags para serialización, validación, etc.
	// Similar a anotaciones en Java (@JsonProperty, etc.)
	BirthDate time.Time `json:"birth_date" validate:"required"`
}

// NewPerson es un constructor idiomático en Go
// En Java sería un constructor estático o builder pattern
func NewPerson(name string, age int, email string) *Person {
	return &Person{
		Name:      name,
		Age:       age,
		email:     email,
		BirthDate: time.Now(),
	}
}

// ============================================================================
// STRUCT EMBEDDING (Composición en lugar de herencia)
// ============================================================================
// En Java usarías herencia: class Employee extends Person
// En Go usas composición (embedding) que es más flexible
// ============================================================================

// Address es un struct embebido
type Address struct {
	Street  string
	City    string
	Country string
}

// Employee tiene un Person embebido (composición)
// Esto significa que Employee "tiene un" Person, no "es un" Person
type Employee struct {
	Person              // Campos de Person están directamente accesibles
	Address             // Campos de Address también
	EmployeeID string
	Salary    float64
	Department string
}

// NewEmployee crea un nuevo empleado
func NewEmployee(name string, age int, email string, employeeID string, salary float64) *Employee {
	return &Employee{
		Person: Person{
			Name:      name,
			Age:       age,
			email:     email,
			BirthDate: time.Now(),
		},
		EmployeeID: employeeID,
		Salary:     salary,
		Department: "Engineering",
	}
}

// ============================================================================
// MÉTODOS EN STRUCTS
// ============================================================================

// GetEmail es un método getter (similar a getEmail() en Java)
// El receptor (receiver) es (p Person) - esto es value receiver
func (p Person) GetEmail() string {
	return p.email
}

// SetEmail es un método setter con pointer receiver
// Usa *Person para modificar el struct original
func (p *Person) SetEmail(email string) {
	p.email = email
}

// String implementa la interfaz Stringer (similar a toString() en Java)
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, Email: %s}", p.Name, p.Age, p.email)
}

// GetFullAddress es un método del Employee que usa campos embebidos
func (e Employee) GetFullAddress() string {
	// Puedes acceder directamente a los campos embebidos
	return fmt.Sprintf("%s, %s, %s", e.Street, e.City, e.Country)
}

// ============================================================================
// COMPARACIÓN DE STRUCTS
// ============================================================================
// En Go, los structs se comparan por valor si todos los campos son comparables
// Similar a equals() en Java, pero automático

func comparePersons() {
	p1 := Person{Name: "John", Age: 30}
	p2 := Person{Name: "John", Age: 30}
	p3 := Person{Name: "Jane", Age: 25}

	fmt.Println("p1 == p2:", p1 == p2) // true
	fmt.Println("p1 == p3:", p1 == p3) // false
}

// ============================================================================
// STRUCTS CON PUNTEROS
// ============================================================================
// Puedes usar punteros para evitar copias grandes y permitir modificación

func demonstratePointers() {
	// Value (copia el struct)
	p1 := Person{Name: "John", Age: 30}
	p1.SetEmail("john@example.com") // Esto NO modifica p1 porque es value receiver

	// Pointer (referencia al struct)
	p2 := &Person{Name: "Jane", Age: 25}
	p2.SetEmail("jane@example.com") // Esto SÍ modifica p2

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

// ============================================================================
// STRUCTS ANÓNIMOS
// ============================================================================
// Útiles para estructuras temporales o configuración

func anonymousStruct() {
	// Struct anónimo (similar a clase anónima en Java, pero más simple)
	config := struct {
		Host string
		Port int
		SSL  bool
	}{
		Host: "localhost",
		Port: 8080,
		SSL:  true,
	}

	fmt.Println("Config:", config)
}

// ============================================================================
// EJEMPLO PRÁCTICO: Sistema de Pagos
// ============================================================================

// PaymentMethod es una interfaz (veremos interfaces después)
type PaymentMethod interface {
	ProcessPayment(amount float64) error
	GetName() string
}

// CreditCard representa una tarjeta de crédito
type CreditCard struct {
	Number     string
	ExpiryDate string
	CVV        string
	CardHolder string
}

// ProcessPayment implementa PaymentMethod para CreditCard
func (c *CreditCard) ProcessPayment(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f for card ending in %s\n", amount, c.Number[len(c.Number)-4:])
	return nil
}

func (c *CreditCard) GetName() string {
	return "Credit Card"
}

// PayPal representa una cuenta PayPal
type PayPal struct {
	Email    string
	Password string // En producción, nunca almacenes passwords en texto plano
}

func (p *PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f for account %s\n", amount, p.Email)
	return nil
}

func (p *PayPal) GetName() string {
	return "PayPal"
}

// ============================================================================
// MAIN - Ejemplos de uso
// ============================================================================

func main() {
	fmt.Println("=== FUNDAMENTOS: TIPOS Y STRUCTS ===\n")

	// 1. Crear structs básicos
	fmt.Println("1. Creando structs básicos:")
	person := NewPerson("John Doe", 30, "john@example.com")
	fmt.Println(person)
	fmt.Println()

	// 2. Struct embedding
	fmt.Println("2. Struct embedding (composición):")
	employee := NewEmployee("Jane Smith", 28, "jane@example.com", "EMP001", 75000.0)
	employee.Street = "123 Main St"
	employee.City = "San Francisco"
	employee.Country = "USA"
	fmt.Printf("Employee: %s\n", employee.Name) // Acceso directo a campos embebidos
	fmt.Printf("Address: %s\n", employee.GetFullAddress())
	fmt.Println()

	// 3. Métodos
	fmt.Println("3. Métodos en structs:")
	fmt.Printf("Email: %s\n", person.GetEmail())
	person.SetEmail("newemail@example.com")
	fmt.Printf("Nuevo email: %s\n", person.GetEmail())
	fmt.Println()

	// 4. Comparación
	fmt.Println("4. Comparación de structs:")
	comparePersons()
	fmt.Println()

	// 5. Punteros
	fmt.Println("5. Uso de punteros:")
	demonstratePointers()
	fmt.Println()

	// 6. Structs anónimos
	fmt.Println("6. Structs anónimos:")
	anonymousStruct()
	fmt.Println()

	// 7. Ejemplo práctico: Sistema de pagos
	fmt.Println("7. Sistema de pagos (preparación para interfaces):")
	creditCard := &CreditCard{
		Number:     "4532-1234-5678-9010",
		ExpiryDate: "12/25",
		CVV:        "123",
		CardHolder: "John Doe",
	}
	paypal := &PayPal{
		Email:    "john@example.com",
		Password: "secret",
	}

	// Procesar pagos (veremos interfaces después)
	creditCard.ProcessPayment(100.50)
	paypal.ProcessPayment(50.25)
	fmt.Println()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

