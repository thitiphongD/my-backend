package ports

import "github.com/thitiphongD/my-backend/internal/core/domain"

// UserRepository defines the interface for user data access
type UserRepository interface {
	// User CRUD operations
	Create(user *domain.User) error
	GetByID(id uint) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
	List() ([]*domain.User, error)

	// Authentication related
	FindByEmailAndPassword(email, password string) (*domain.User, error)
}
