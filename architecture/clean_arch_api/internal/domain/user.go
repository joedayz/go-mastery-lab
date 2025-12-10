package domain

// ============================================================================
// DOMAIN LAYER - Entidades puras sin dependencias externas
// ============================================================================
// En Clean Architecture, el dominio es el núcleo
// No debe depender de frameworks, bases de datos, etc.
// ============================================================================

// User representa una entidad de usuario en el dominio
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"` // No serializar password
}

// Validate valida los datos del usuario
func (u *User) Validate() error {
	if u.Email == "" {
		return ErrInvalidEmail
	}
	if u.Name == "" {
		return ErrInvalidName
	}
	return nil
}

// Errores del dominio
var (
	ErrInvalidEmail = &DomainError{Message: "invalid email"}
	ErrInvalidName  = &DomainError{Message: "invalid name"}
	ErrUserNotFound = &DomainError{Message: "user not found"}
)

type DomainError struct {
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

