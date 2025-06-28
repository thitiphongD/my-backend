package auth

import (
	"errors"

	"github.com/thitiphongD/my-backend/internal/database"
	"github.com/thitiphongD/my-backend/internal/models"
	"github.com/thitiphongD/my-backend/internal/utils"
	"gorm.io/gorm"
)

// AuthService handles authentication related operations
type AuthService struct {
	db *gorm.DB
}

// NewAuthService creates a new auth service instance
func NewAuthService() *AuthService {
	return &AuthService{
		db: database.GetDB(),
	}
}

// Register creates a new user account
func (s *AuthService) Register(req models.RegisterRequest) (*models.AuthResponse, error) {
	// Check if user already exists
	var existingUser models.User
	result := s.db.Where("email = ?", req.Email).First(&existingUser)
	if result.Error == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create new user
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	result = s.db.Create(&user)
	if result.Error != nil {
		return nil, errors.New("failed to create user")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &models.AuthResponse{
		Token: token,
		User:  &user,
	}, nil
}

// Login authenticates a user
func (s *AuthService) Login(req models.LoginRequest) (*models.AuthResponse, error) {
	// Find user by email
	var user models.User
	result := s.db.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
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

	return &models.AuthResponse{
		Token: token,
		User:  &user,
	}, nil
}

// GetUserByID retrieves a user by ID
func (s *AuthService) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	result := s.db.First(&user, userID)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
