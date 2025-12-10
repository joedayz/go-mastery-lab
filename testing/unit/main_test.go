package main

import (
	"testing"
)

// ============================================================================
// TESTING AVANZADO EN GO
// ============================================================================
// Go tiene un framework de testing integrado muy poderoso
// Similar a JUnit en Java, pero más simple y directo
// ============================================================================

// ============================================================================
// TEST BÁSICO
// ============================================================================

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// ============================================================================
// TABLE-DRIVEN TESTS (BEST PRACTICE EN GO)
// ============================================================================
// Los table-driven tests son el patrón más común en Go
// Permiten probar múltiples casos fácilmente

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"zero", 0, 5, 5},
		{"mixed", -2, 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// ============================================================================
// MOCKS CON INTERFACES
// ============================================================================
// En Go, los mocks se hacen fácilmente con interfaces
// No necesitas librerías complejas como Mockito en Java

type Calculator interface {
	Add(a, b int) int
	Multiply(a, b int) int
}

type RealCalculator struct{}

func (c *RealCalculator) Add(a, b int) int {
	return a + b
}

func (c *RealCalculator) Multiply(a, b int) int {
	return a * b
}

// MockCalculator para testing
type MockCalculator struct {
	AddFunc      func(int, int) int
	MultiplyFunc func(int, int) int
}

func (m *MockCalculator) Add(a, b int) int {
	if m.AddFunc != nil {
		return m.AddFunc(a, b)
	}
	return 0
}

func (m *MockCalculator) Multiply(a, b int) int {
	if m.MultiplyFunc != nil {
		return m.MultiplyFunc(a, b)
	}
	return 0
}

func TestWithMock(t *testing.T) {
	mockCalc := &MockCalculator{
		AddFunc: func(a, b int) int {
			return 999 // Valor mock
		},
	}

	result := mockCalc.Add(2, 3)
	if result != 999 {
		t.Errorf("Expected 999, got %d", result)
	}
}

// ============================================================================
// TESTING CON HELPERS
// ============================================================================

func assertEqual(t *testing.T, got, want int) {
	t.Helper() // Marca esta función como helper
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestWithHelper(t *testing.T) {
	assertEqual(t, Add(2, 3), 5)
	assertEqual(t, Add(0, 0), 0)
}

// ============================================================================
// SUBTESTS
// ============================================================================

func TestSubtests(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		if Add(2, 3) != 5 {
			t.Error("addition failed")
		}
	})

	t.Run("subtraction", func(t *testing.T) {
		// Implementar resta
	})
}

// ============================================================================
// BENCHMARKS
// ============================================================================

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// ============================================================================
// EJEMPLO PRÁCTICO: TESTING DE REPOSITORIO
// ============================================================================

type UserRepository interface {
	GetUser(id int) (*User, error)
}

type User struct {
	ID   int
	Name string
}

type MockUserRepository struct {
	users map[int]*User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[int]*User),
	}
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, &NotFoundError{ID: id}
	}
	return user, nil
}

func (m *MockUserRepository) SetUser(user *User) {
	m.users[user.ID] = user
}

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return "user not found"
}

func TestUserRepository(t *testing.T) {
	repo := NewMockUserRepository()
	repo.SetUser(&User{ID: 1, Name: "John"})

	user, err := repo.GetUser(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Name != "John" {
		t.Errorf("expected John, got %s", user.Name)
	}

	_, err = repo.GetUser(999)
	if err == nil {
		t.Error("expected error for non-existent user")
	}
}

