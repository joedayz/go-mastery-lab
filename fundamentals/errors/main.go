package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// ============================================================================
// MANEJO DE ERRORES EN GO
// ============================================================================
// En Go, los errores son valores, no excepciones
// No hay try-catch como en Java
// La convención es retornar error como último valor
// ============================================================================

// ============================================================================
// ERRORES BÁSICOS
// ============================================================================

// Crear errores simples
var (
	ErrNotFound      = errors.New("resource not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrUnauthorized  = errors.New("unauthorized")
)

// Crear errores con formato
func createFormattedError(id int) error {
	return fmt.Errorf("user with id %d not found", id)
}

// ============================================================================
// ERRORES PERSONALIZADOS
// ============================================================================
// Puedes crear tipos de error personalizados

// ValidationError es un error personalizado
type ValidationError struct {
	Field   string
	Message string
}

// Error implementa la interfaz error
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// ============================================================================
// WRAPPING ERRORS (Go 1.13+)
// ============================================================================
// Wrapping permite agregar contexto a errores sin perder el error original
// Similar a causa en Java (cause), pero más idiomático

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		// Wrap el error agregando contexto
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil && err != io.EOF {
		// Wrap el error agregando contexto
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	return nil
}

// ============================================================================
// ERRORS.IS() - VERIFICAR ERRORES ESPECÍFICOS
// ============================================================================
// errors.Is verifica si un error es o contiene un error específico
// Útil para comparar con errores predefinidos

func demonstrateErrorsIs() {
	err := readFile("nonexistent.txt")

	// Verificar si el error es específicamente os.ErrNotExist
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	}

	// Verificar con nuestro error personalizado
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Resource not found")
	}
}

// ============================================================================
// ERRORS.AS() - EXTRAER ERRORES TIPO ESPECÍFICO
// ============================================================================
// errors.As extrae un error de un tipo específico de la cadena de errores
// Similar a instanceof en Java, pero para errores

func demonstrateErrorsAs() {
	err := &ValidationError{
		Field:   "email",
		Message: "invalid format",
	}

	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("Validation error on field: %s\n", validationErr.Field)
		fmt.Printf("Message: %s\n", validationErr.Message)
	}
}

// ============================================================================
// ERRORS.JOIN() - COMBINAR MÚLTIPLES ERRORES
// ============================================================================
// errors.Join combina múltiples errores en uno solo
// Útil para validaciones que pueden tener múltiples errores

func validateUser(name, email string) error {
	var errs []error

	if name == "" {
		errs = append(errs, &ValidationError{Field: "name", Message: "cannot be empty"})
	}
	if email == "" {
		errs = append(errs, &ValidationError{Field: "email", Message: "cannot be empty"})
	}
	if len(email) > 0 && !contains(email, "@") {
		errs = append(errs, &ValidationError{Field: "email", Message: "invalid format"})
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// ============================================================================
// PATRÓN: SENTINEL ERRORS
// ============================================================================
// Errores predefinidos que se pueden comparar directamente
// Útil para errores conocidos que el llamador puede manejar

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserExists      = errors.New("user already exists")
)

type User struct {
	ID       int
	Username string
	Email    string
}

type UserRepository struct {
	users map[int]*User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int]*User),
	}
}

func (r *UserRepository) CreateUser(id int, username, email string) error {
	if _, exists := r.users[id]; exists {
		return ErrUserExists
	}
	r.users[id] = &User{ID: id, Username: username, Email: email}
	return nil
}

func (r *UserRepository) GetUser(id int) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("failed to get user: %w", ErrUserNotFound)
	}
	return user, nil
}

func (r *UserRepository) Authenticate(username, password string) (*User, error) {
	for _, user := range r.users {
		if user.Username == username {
			// Simulación: password correcto es "password123"
			if password != "password123" {
				return nil, fmt.Errorf("authentication failed: %w", ErrInvalidPassword)
			}
			return user, nil
		}
	}
	return nil, fmt.Errorf("authentication failed: %w", ErrUserNotFound)
}

// ============================================================================
// MANEJO DE ERRORES EN FUNCIONES
// ============================================================================

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero: cannot divide %f by %f", a, b)
	}
	return a / b, nil
}

func processUser(userID int) error {
	repo := NewUserRepository()

	// Crear usuario
	if err := repo.CreateUser(1, "john", "john@example.com"); err != nil {
		if errors.Is(err, ErrUserExists) {
			fmt.Println("User already exists, continuing...")
		} else {
			return fmt.Errorf("failed to create user: %w", err)
		}
	}

	// Obtener usuario
	user, err := repo.GetUser(userID)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return fmt.Errorf("user %d not found", userID)
		}
		return err
	}

	fmt.Printf("Found user: %+v\n", user)
	return nil
}

// ============================================================================
// EJEMPLO COMPLETO: MANEJO DE ERRORES EN CAPAS
// ============================================================================

// Service layer
type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(id int, username, email, password string) error {
	// Validar entrada
	if err := validateUser(username, email); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Crear usuario
	if err := s.repo.CreateUser(id, username, email); err != nil {
		if errors.Is(err, ErrUserExists) {
			return fmt.Errorf("registration failed: user already exists")
		}
		return fmt.Errorf("registration failed: %w", err)
	}

	return nil
}

func (s *UserService) Login(username, password string) (*User, error) {
	user, err := s.repo.Authenticate(username, password)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, fmt.Errorf("login failed: user not found")
		}
		if errors.Is(err, ErrInvalidPassword) {
			return nil, fmt.Errorf("login failed: invalid password")
		}
		return nil, fmt.Errorf("login failed: %w", err)
	}
	return user, nil
}

// ============================================================================
// MAIN - Ejemplos de uso
// ============================================================================

func main() {
	fmt.Println("=== FUNDAMENTOS: MANEJO DE ERRORES ===\n")

	// 1. Errores básicos
	fmt.Println("1. Errores básicos:")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error (expected): %v\n", err)
	}
	fmt.Println()

	// 2. Wrapping errors
	fmt.Println("2. Wrapping errors:")
	if err := readFile("nonexistent.txt"); err != nil {
		fmt.Printf("Error: %v\n", err)
		// Verificar el error original
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("  -> Original error is os.ErrNotExist")
		}
	}
	fmt.Println()

	// 3. errors.Is()
	fmt.Println("3. errors.Is():")
	demonstrateErrorsIs()
	fmt.Println()

	// 4. errors.As()
	fmt.Println("4. errors.As():")
	demonstrateErrorsAs()
	fmt.Println()

	// 5. errors.Join()
	fmt.Println("5. errors.Join() - Validación múltiple:")
	if err := validateUser("", ""); err != nil {
		fmt.Printf("Validation errors: %v\n", err)
		// errors.Join crea un error que contiene múltiples errores
		if joinedErr, ok := err.(interface{ Unwrap() []error }); ok {
			fmt.Println("  Individual errors:")
			for _, e := range joinedErr.Unwrap() {
				fmt.Printf("    - %v\n", e)
			}
		}
	}
	fmt.Println()

	// 6. Sentinel errors
	fmt.Println("6. Sentinel errors:")
	repo := NewUserRepository()
	if err := repo.CreateUser(1, "john", "john@example.com"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	if err := repo.CreateUser(1, "jane", "jane@example.com"); err != nil {
		if errors.Is(err, ErrUserExists) {
			fmt.Printf("Expected error: %v\n", err)
		}
	}
	fmt.Println()

	// 7. Manejo de errores en capas
	fmt.Println("7. Manejo de errores en capas:")
	service := NewUserService(repo)
	if err := service.RegisterUser(2, "alice", "alice@example.com", "password123"); err != nil {
		fmt.Printf("Registration error: %v\n", err)
	} else {
		fmt.Println("User registered successfully")
	}

	user, err := service.Login("alice", "wrongpassword")
	if err != nil {
		fmt.Printf("Login error: %v\n", err)
		if errors.Is(err, ErrInvalidPassword) {
			fmt.Println("  -> Password was incorrect")
		}
	} else {
		fmt.Printf("Logged in user: %+v\n", user)
	}
	fmt.Println()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

