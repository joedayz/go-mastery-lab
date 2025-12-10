package infrastructure

import (
	"fmt"
	"sync"
	"github.com/josediaz/go-mastery-lab/architecture/clean_arch_api/internal/domain"
	"github.com/josediaz/go-mastery-lab/architecture/clean_arch_api/internal/repository"
)

// ============================================================================
// INFRASTRUCTURE LAYER - Implementaciones concretas
// ============================================================================
// Esta es la implementación concreta del repositorio
// En producción, aquí estaría la conexión a la base de datos
// ============================================================================

// MemoryUserRepository es una implementación en memoria del UserRepository
type MemoryUserRepository struct {
	mu    sync.RWMutex
	users map[int]*domain.User
	nextID int
}

func NewMemoryUserRepository() repository.UserRepository {
	return &MemoryUserRepository{
		users: make(map[int]*domain.User),
		nextID: 1,
	}
}

func (r *MemoryUserRepository) Create(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) GetByID(id int) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

func (r *MemoryUserRepository) GetByEmail(email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

func (r *MemoryUserRepository) Update(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return domain.ErrUserNotFound
	}
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return domain.ErrUserNotFound
	}
	delete(r.users, id)
	return nil
}

