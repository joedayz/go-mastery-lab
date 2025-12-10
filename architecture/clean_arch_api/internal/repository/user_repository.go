package repository

import "github.com/josediaz/go-mastery-lab/architecture/clean_arch_api/internal/domain"

// ============================================================================
// REPOSITORY LAYER - Interfaces para acceso a datos
// ============================================================================
// Las interfaces están en el dominio o en una capa de repositorio
// Las implementaciones concretas están en infrastructure
// ============================================================================

// UserRepository define el contrato para acceso a usuarios
type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id int) error
}

