package repositories

import (
	"errors"

	"github.com/thitiphongD/my-backend/internal/core/domain"
	"github.com/thitiphongD/my-backend/internal/core/ports"
	"gorm.io/gorm"
)

// mangaRepository implements the MangaRepository interface
type mangaRepository struct {
	db *gorm.DB
}

// NewMangaRepository creates a new manga repository instance
func NewMangaRepository(db *gorm.DB) ports.MangaRepository {
	return &mangaRepository{
		db: db,
	}
}

// Create creates a new manga in the database
func (r *mangaRepository) Create(manga *domain.Manga) error {
	if err := r.db.Create(manga).Error; err != nil {
		return errors.New("failed to create manga")
	}
	return nil
}

// GetByID retrieves a manga by ID
func (r *mangaRepository) GetByID(id uint) (*domain.Manga, error) {
	var manga domain.Manga
	if err := r.db.First(&manga, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("manga not found")
		}
		return nil, errors.New("failed to get manga")
	}
	return &manga, nil
}

// GetByUserID retrieves mangas by user ID
func (r *mangaRepository) GetByUserID(userID uint) ([]*domain.Manga, error) {
	var mangas []*domain.Manga
	if err := r.db.Where("user_created = ?", userID).Find(&mangas).Error; err != nil {
		return nil, errors.New("failed to get user mangas")
	}
	return mangas, nil
}

// List retrieves all mangas from the database
func (r *mangaRepository) List() ([]*domain.Manga, error) {
	var mangas []*domain.Manga
	if err := r.db.Find(&mangas).Error; err != nil {
		return nil, errors.New("failed to get mangas")
	}
	return mangas, nil
}

// Update updates a manga in the database
func (r *mangaRepository) Update(manga *domain.Manga) error {
	if err := r.db.Save(manga).Error; err != nil {
		return errors.New("failed to update manga")
	}
	return nil
}

// Delete soft deletes a manga from the database
func (r *mangaRepository) Delete(id uint) error {
	if err := r.db.Delete(&domain.Manga{}, id).Error; err != nil {
		return errors.New("failed to delete manga")
	}
	return nil
}

// GetActiveMangas retrieves all active mangas
func (r *mangaRepository) GetActiveMangas() ([]*domain.Manga, error) {
	var mangas []*domain.Manga
	if err := r.db.Where("is_active = ?", true).Find(&mangas).Error; err != nil {
		return nil, errors.New("failed to get active mangas")
	}
	return mangas, nil
}

// GetMangasByPriceRange retrieves mangas within price range
func (r *mangaRepository) GetMangasByPriceRange(min, max float64) ([]*domain.Manga, error) {
	var mangas []*domain.Manga
	if err := r.db.Where("price BETWEEN ? AND ?", min, max).Find(&mangas).Error; err != nil {
		return nil, errors.New("failed to get mangas by price range")
	}
	return mangas, nil
}

// ListPaginated retrieves mangas with pagination
func (r *mangaRepository) ListPaginated(pagination *domain.PaginationRequest) ([]*domain.Manga, int64, error) {
	var mangas []*domain.Manga
	var total int64

	// Count total records
	if err := r.db.Model(&domain.Manga{}).Count(&total).Error; err != nil {
		return nil, 0, errors.New("failed to count mangas")
	}

	// Get paginated results
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()

	if err := r.db.Offset(offset).Limit(limit).Find(&mangas).Error; err != nil {
		return nil, 0, errors.New("failed to get paginated mangas")
	}

	return mangas, total, nil
}

// GetActiveMangasPaginated retrieves active mangas with pagination
func (r *mangaRepository) GetActiveMangasPaginated(pagination *domain.PaginationRequest) ([]*domain.Manga, int64, error) {
	var mangas []*domain.Manga
	var total int64

	// Count total active records
	if err := r.db.Model(&domain.Manga{}).Where("is_active = ?", true).Count(&total).Error; err != nil {
		return nil, 0, errors.New("failed to count active mangas")
	}

	// Get paginated results
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()

	if err := r.db.Where("is_active = ?", true).Offset(offset).Limit(limit).Find(&mangas).Error; err != nil {
		return nil, 0, errors.New("failed to get paginated active mangas")
	}

	return mangas, total, nil
}

// GetMangasByUserIDPaginated retrieves mangas by user ID with pagination
func (r *mangaRepository) GetMangasByUserIDPaginated(userID uint, pagination *domain.PaginationRequest) ([]*domain.Manga, int64, error) {
	var mangas []*domain.Manga
	var total int64

	// Count total user records
	if err := r.db.Model(&domain.Manga{}).Where("user_created = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, errors.New("failed to count user mangas")
	}

	// Get paginated results
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()

	if err := r.db.Where("user_created = ?", userID).Offset(offset).Limit(limit).Find(&mangas).Error; err != nil {
		return nil, 0, errors.New("failed to get paginated user mangas")
	}

	return mangas, total, nil
}

// GetMangasByPriceRangePaginated retrieves mangas within price range with pagination
func (r *mangaRepository) GetMangasByPriceRangePaginated(min, max float64, pagination *domain.PaginationRequest) ([]*domain.Manga, int64, error) {
	var mangas []*domain.Manga
	var total int64

	// Count total records in price range
	if err := r.db.Model(&domain.Manga{}).Where("price BETWEEN ? AND ?", min, max).Count(&total).Error; err != nil {
		return nil, 0, errors.New("failed to count mangas by price range")
	}

	// Get paginated results
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()

	if err := r.db.Where("price BETWEEN ? AND ?", min, max).Offset(offset).Limit(limit).Find(&mangas).Error; err != nil {
		return nil, 0, errors.New("failed to get paginated mangas by price range")
	}

	return mangas, total, nil
}
