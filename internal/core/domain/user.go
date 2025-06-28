package domain

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user entity in the domain
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"` // "-" excludes from JSON serialization
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// IsValid checks if the user has valid data
func (u *User) IsValid() bool {
	return u.Name != "" && u.Email != "" && u.Password != ""
}

// Sanitize removes sensitive data from user before returning
func (u *User) Sanitize() *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
