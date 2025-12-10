package usecase

import (
	"context"
	"github.com/josediaz/go-mastery-lab/architecture/clean_arch_api/internal/domain"
	"github.com/josediaz/go-mastery-lab/architecture/clean_arch_api/internal/repository"
)

// ============================================================================
// USECASE LAYER - Lógica de negocio
// ============================================================================
// Los casos de uso orquestan el flujo de la aplicación
// Dependen de interfaces (repositorios), no de implementaciones
// ============================================================================

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

// CreateUser crea un nuevo usuario
func (uc *UserUsecase) CreateUser(ctx context.Context, email, name, password string) (*domain.User, error) {
	// Validar entrada
	user := &domain.User{
		Email:    email,
		Name:     name,
		Password: password,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Verificar si el usuario ya existe
	existing, _ := uc.userRepo.GetByEmail(email)
	if existing != nil {
		return nil, domain.ErrUserNotFound // En producción, usaría otro error
	}

	// Crear usuario
	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser obtiene un usuario por ID
func (uc *UserUsecase) GetUser(ctx context.Context, id int) (*domain.User, error) {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

