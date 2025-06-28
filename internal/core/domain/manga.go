package domain

import (
	"time"

	"gorm.io/gorm"
)

// Manga represents the manga entity in the domain
type Manga struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name" gorm:"not null"`
	Price       float64        `json:"price" gorm:"not null"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	UserCreated uint           `json:"user_created" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// IsValid checks if the manga has valid data
func (m *Manga) IsValid() bool {
	return m.Name != "" && m.Price >= 0 && m.UserCreated > 0
}

// Sanitize removes sensitive data from manga before returning
func (m *Manga) Sanitize() *Manga {
	return &Manga{
		ID:          m.ID,
		Name:        m.Name,
		Price:       m.Price,
		IsActive:    m.IsActive,
		UserCreated: m.UserCreated,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
