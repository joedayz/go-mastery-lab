package main

import (
	"fmt"
	"math"
)

// ============================================================================
// INTERFACES IMPLÍCITAS (DUCK TYPING)
// ============================================================================
// En Go, las interfaces son implícitas. Si un tipo implementa todos los
// métodos de una interfaz, automáticamente la implementa.
// No necesitas declarar "implements" como en Java.
//
// "If it walks like a duck and quacks like a duck, it's a duck"
// ============================================================================

// ============================================================================
// INTERFAZ BÁSICA
// ============================================================================

// Shape define una forma geométrica
// En Java sería: interface Shape { double area(); double perimeter(); }
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ============================================================================
// IMPLEMENTACIONES DE LA INTERFAZ
// ============================================================================
// Nota: No hay palabra clave "implements" en Go
// Si Circle tiene métodos Area() y Perimeter(), automáticamente implementa Shape

// Circle representa un círculo
type Circle struct {
	Radius float64
}

// Area calcula el área del círculo
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calcula el perímetro del círculo
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle representa un rectángulo
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ============================================================================
// INTERFACES PEQUEÑAS Y ESPECÍFICAS (BEST PRACTICE)
// ============================================================================
// En Go, es mejor tener interfaces pequeñas y específicas
// Similar al principio de segregación de interfaces en SOLID

// Writer es una interfaz del paquete io (estándar de Go)
// type Writer interface {
//     Write(p []byte) (n int, err error)
// }

// Reader es otra interfaz del paquete io
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }

// ============================================================================
// EJEMPLO PRÁCTICO: SISTEMA DE PAGOS CON MÚLTIPLES PROVEEDORES
// ============================================================================
// Este es el ejemplo sugerido en el temario: una librería de pagos
// con varios proveedores usando interfaces

// PaymentProvider define el contrato para procesadores de pago
// Similar a una interfaz en Java: interface PaymentProvider { ... }
type PaymentProvider interface {
	// ProcessPayment procesa un pago y retorna un error si falla
	ProcessPayment(amount float64, currency string) error

	// GetName retorna el nombre del proveedor
	GetName() string

	// IsAvailable verifica si el proveedor está disponible
	IsAvailable() bool
}

// ============================================================================
// IMPLEMENTACIONES DE PAYMENT PROVIDER
// ============================================================================

// StripeProvider implementa PaymentProvider para Stripe
// En Java: class StripeProvider implements PaymentProvider
type StripeProvider struct {
	APIKey    string
	IsEnabled bool
}

func NewStripeProvider(apiKey string) *StripeProvider {
	return &StripeProvider{
		APIKey:    apiKey,
		IsEnabled: true,
	}
}

func (s *StripeProvider) ProcessPayment(amount float64, currency string) error {
	if !s.IsEnabled {
		return fmt.Errorf("stripe provider is disabled")
	}
	fmt.Printf("[Stripe] Processing payment: %.2f %s\n", amount, currency)
	fmt.Printf("[Stripe] Using API Key: %s...\n", s.APIKey[:10])
	// Aquí iría la lógica real de Stripe
	return nil
}

func (s *StripeProvider) GetName() string {
	return "Stripe"
}

func (s *StripeProvider) IsAvailable() bool {
	return s.IsEnabled
}

// PayPalProvider implementa PaymentProvider para PayPal
type PayPalProvider struct {
	ClientID     string
	ClientSecret string
	IsEnabled    bool
}

func NewPayPalProvider(clientID, clientSecret string) *PayPalProvider {
	return &PayPalProvider{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		IsEnabled:    true,
	}
}

func (p *PayPalProvider) ProcessPayment(amount float64, currency string) error {
	if !p.IsAvailable() {
		return fmt.Errorf("paypal provider is not available")
	}
	fmt.Printf("[PayPal] Processing payment: %.2f %s\n", amount, currency)
	fmt.Printf("[PayPal] Using Client ID: %s\n", p.ClientID)
	// Aquí iría la lógica real de PayPal
	return nil
}

func (p *PayPalProvider) GetName() string {
	return "PayPal"
}

func (p *PayPalProvider) IsAvailable() bool {
	return p.IsEnabled
}

// BankTransferProvider implementa PaymentProvider para transferencias bancarias
type BankTransferProvider struct {
	BankName string
	Account  string
	IsEnabled bool
}

func NewBankTransferProvider(bankName, account string) *BankTransferProvider {
	return &BankTransferProvider{
		BankName:  bankName,
		Account:   account,
		IsEnabled: true,
	}
}

func (b *BankTransferProvider) ProcessPayment(amount float64, currency string) error {
	if !b.IsAvailable() {
		return fmt.Errorf("bank transfer provider is not available")
	}
	fmt.Printf("[Bank Transfer] Processing payment: %.2f %s\n", amount, currency)
	fmt.Printf("[Bank Transfer] Bank: %s, Account: %s\n", b.BankName, b.Account)
	return nil
}

func (b *BankTransferProvider) GetName() string {
	return "Bank Transfer"
}

func (b *BankTransferProvider) IsAvailable() bool {
	return b.IsEnabled
}

// ============================================================================
// CLIENTE DE PAGO QUE USA LA INTERFAZ
// ============================================================================
// PaymentService es agnóstico del proveedor específico
// Solo conoce la interfaz PaymentProvider
// Esto es inversión de dependencias (Dependency Inversion Principle)

type PaymentService struct {
	providers []PaymentProvider
}

func NewPaymentService(providers ...PaymentProvider) *PaymentService {
	return &PaymentService{
		providers: providers,
	}
}

// ProcessPaymentWithProvider procesa un pago con un proveedor específico
func (ps *PaymentService) ProcessPaymentWithProvider(providerName string, amount float64, currency string) error {
	for _, provider := range ps.providers {
		if provider.GetName() == providerName && provider.IsAvailable() {
			return provider.ProcessPayment(amount, currency)
		}
	}
	return fmt.Errorf("provider %s not found or not available", providerName)
}

// ProcessPaymentWithFirstAvailable procesa con el primer proveedor disponible
func (ps *PaymentService) ProcessPaymentWithFirstAvailable(amount float64, currency string) error {
	for _, provider := range ps.providers {
		if provider.IsAvailable() {
			fmt.Printf("Using provider: %s\n", provider.GetName())
			return provider.ProcessPayment(amount, currency)
		}
	}
	return fmt.Errorf("no payment provider available")
}

// ListAvailableProviders lista todos los proveedores disponibles
func (ps *PaymentService) ListAvailableProviders() []string {
	var available []string
	for _, provider := range ps.providers {
		if provider.IsAvailable() {
			available = append(available, provider.GetName())
		}
	}
	return available
}

// ============================================================================
// TYPE ASSERTIONS Y TYPE SWITCHES
// ============================================================================
// A veces necesitas verificar o convertir tipos en interfaces

func demonstrateTypeAssertion(provider PaymentProvider) {
	// Type assertion: verifica si provider es específicamente un *StripeProvider
	stripe, ok := provider.(*StripeProvider)
	if ok {
		fmt.Printf("This is a Stripe provider with API Key: %s\n", stripe.APIKey[:10])
	} else {
		fmt.Println("This is not a Stripe provider")
	}

	// Type switch: similar a switch pero para tipos
	switch v := provider.(type) {
	case *StripeProvider:
		fmt.Printf("Stripe provider detected: %s\n", v.APIKey[:10])
	case *PayPalProvider:
		fmt.Printf("PayPal provider detected: %s\n", v.ClientID)
	case *BankTransferProvider:
		fmt.Printf("Bank Transfer provider detected: %s\n", v.BankName)
	default:
		fmt.Printf("Unknown provider type: %T\n", v)
	}
}

// ============================================================================
// INTERFACES VACÍAS (interface{})
// ============================================================================
// interface{} (o any desde Go 1.18) acepta cualquier tipo
// Similar a Object en Java, pero más flexible

func acceptAnything(value interface{}) {
	fmt.Printf("Received value of type %T: %v\n", value, value)
}

// ============================================================================
// MAIN - Ejemplos de uso
// ============================================================================

func main() {
	fmt.Println("=== FUNDAMENTOS: INTERFACES IMPLÍCITAS ===\n")

	// 1. Interfaces básicas con formas geométricas
	fmt.Println("1. Interfaces básicas (Shapes):")
	shapes := []Shape{
		Circle{Radius: 5.0},
		Rectangle{Width: 4.0, Height: 6.0},
		Circle{Radius: 3.0},
	}

	for _, shape := range shapes {
		fmt.Printf("Shape: Area=%.2f, Perimeter=%.2f\n", shape.Area(), shape.Perimeter())
	}
	fmt.Println()

	// 2. Sistema de pagos con múltiples proveedores
	fmt.Println("2. Sistema de pagos con múltiples proveedores:")
	
	// Crear proveedores
	stripe := NewStripeProvider("sk_live_1234567890abcdef")
	paypal := NewPayPalProvider("paypal_client_id", "paypal_secret")
	bankTransfer := NewBankTransferProvider("Chase Bank", "ACC123456")

	// Crear servicio de pagos con todos los proveedores
	paymentService := NewPaymentService(stripe, paypal, bankTransfer)

	// Listar proveedores disponibles
	fmt.Println("Available providers:", paymentService.ListAvailableProviders())
	fmt.Println()

	// Procesar pago con proveedor específico
	fmt.Println("3. Procesando pago con Stripe:")
	if err := paymentService.ProcessPaymentWithProvider("Stripe", 99.99, "USD"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()

	// Procesar pago con primer proveedor disponible
	fmt.Println("4. Procesando pago con primer proveedor disponible:")
	if err := paymentService.ProcessPaymentWithFirstAvailable(50.00, "EUR"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()

	// 5. Type assertions y type switches
	fmt.Println("5. Type assertions y type switches:")
	demonstrateTypeAssertion(stripe)
	fmt.Println()
	demonstrateTypeAssertion(paypal)
	fmt.Println()

	// 6. Interfaces vacías
	fmt.Println("6. Interfaces vacías (interface{} / any):")
	acceptAnything(42)
	acceptAnything("hello")
	acceptAnything([]int{1, 2, 3})
	acceptAnything(stripe)
	fmt.Println()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

