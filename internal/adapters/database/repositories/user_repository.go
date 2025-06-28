package repositories

import (
	"errors"

	"github.com/thitiphongD/my-backend/internal/core/domain"
	"github.com/thitiphongD/my-backend/internal/core/ports"
	"gorm.io/gorm"
)

// userRepository implements the UserRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create creates a new user in the database
func (r *userRepository) Create(user *domain.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to get user")
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to get user")
	}
	return &user, nil
}

// Update updates a user in the database
func (r *userRepository) Update(user *domain.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.New("failed to update user")
	}
	return nil
}

// Delete soft deletes a user from the database
func (r *userRepository) Delete(id uint) error {
	if err := r.db.Delete(&domain.User{}, id).Error; err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}

// List retrieves all users from the database
func (r *userRepository) List() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, errors.New("failed to get users")
	}
	return users, nil
}

// FindByEmailAndPassword finds a user by email and password (for login)
func (r *userRepository) FindByEmailAndPassword(email, password string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, errors.New("failed to authenticate user")
	}
	return &user, nil
}
