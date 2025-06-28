package ports

import "github.com/thitiphongD/my-backend/internal/core/domain"

// AuthService defines the interface for authentication operations
type AuthService interface {
	Register(req *domain.RegisterRequest) (*domain.AuthResponse, error)
	Login(req *domain.LoginRequest) (*domain.AuthResponse, error)
	GetUserByID(userID uint) (*domain.User, error)
	ValidateToken(token string) (*domain.User, error)
}

// UserService defines the interface for user operations
type UserService interface {
	CreateUser(req *domain.CreateUserRequest) (*domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	UpdateUser(id uint, req *domain.CreateUserRequest) (*domain.User, error)
	DeleteUser(id uint) error
}
