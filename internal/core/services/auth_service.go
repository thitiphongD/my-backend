package services

import (
	"errors"

	"github.com/thitiphongD/my-backend/internal/core/domain"
	"github.com/thitiphongD/my-backend/internal/core/ports"
	"github.com/thitiphongD/my-backend/internal/utils"
)

// authService implements the AuthService interface
type authService struct {
	userRepo ports.UserRepository
}

// NewAuthService creates a new auth service instance
func NewAuthService(userRepo ports.UserRepository) ports.AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// Register creates a new user account
func (s *authService) Register(req *domain.RegisterRequest) (*domain.AuthResponse, error) {
	// Check if user already exists
	_, err := s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create new user
	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if !user.IsValid() {
		return nil, errors.New("invalid user data")
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &domain.AuthResponse{
		Token: token,
		User:  user.Sanitize(),
	}, nil
}

// Login authenticates a user
func (s *authService) Login(req *domain.LoginRequest) (*domain.AuthResponse, error) {
	// Find user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &domain.AuthResponse{
		Token: token,
		User:  user.Sanitize(),
	}, nil
}

// GetUserByID retrieves a user by ID
func (s *authService) GetUserByID(userID uint) (*domain.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user.Sanitize(), nil
}

// ValidateToken validates a JWT token and returns the user
func (s *authService) ValidateToken(token string) (*domain.User, error) {
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	user, err := s.userRepo.GetByID(claims.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user.Sanitize(), nil
}
