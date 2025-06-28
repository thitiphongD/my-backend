package services

import (
	"errors"

	"github.com/thitiphongD/my-backend/internal/core/domain"
	"github.com/thitiphongD/my-backend/internal/core/ports"
)

// userService implements the UserService interface
type userService struct {
	userRepo ports.UserRepository
}

// NewUserService creates a new user service instance
func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (s *userService) CreateUser(req *domain.CreateUserRequest) (*domain.User, error) {
	// Check if user already exists
	_, err := s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Create new user
	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user.Sanitize(), nil
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user.Sanitize(), nil
}

// GetUsers retrieves all users
func (s *userService) GetUsers() ([]*domain.User, error) {
	users, err := s.userRepo.List()
	if err != nil {
		return nil, err
	}

	// Sanitize all users
	sanitizedUsers := make([]*domain.User, len(users))
	for i, user := range users {
		sanitizedUsers[i] = user.Sanitize()
	}

	return sanitizedUsers, nil
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(id uint, req *domain.CreateUserRequest) (*domain.User, error) {
	// Get existing user
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update user fields
	user.Name = req.Name
	user.Email = req.Email

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user.Sanitize(), nil
}

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(id uint) error {
	// Check if user exists
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(id)
}
