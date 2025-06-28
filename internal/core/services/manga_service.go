package services

import (
	"errors"

	"github.com/thitiphongD/my-backend/internal/core/domain"
	"github.com/thitiphongD/my-backend/internal/core/ports"
)

// mangaService implements the MangaService interface
type mangaService struct {
	mangaRepo ports.MangaRepository
}

// NewMangaService creates a new manga service instance
func NewMangaService(mangaRepo ports.MangaRepository) ports.MangaService {
	return &mangaService{
		mangaRepo: mangaRepo,
	}
}

// CreateManga creates a new manga
func (s *mangaService) CreateManga(req *domain.CreateMangaRequest, userID uint) (*domain.Manga, error) {
	manga := &domain.Manga{
		Name:        req.Name,
		Price:       req.Price,
		IsActive:    req.IsActive,
		UserCreated: userID,
	}

	if !manga.IsValid() {
		return nil, errors.New("invalid manga data")
	}

	if err := s.mangaRepo.Create(manga); err != nil {
		return nil, err
	}

	return manga.Sanitize(), nil
}

// GetMangaByID retrieves a manga by ID
func (s *mangaService) GetMangaByID(id uint) (*domain.Manga, error) {
	manga, err := s.mangaRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return manga.Sanitize(), nil
}

// GetMangas retrieves all mangas
func (s *mangaService) GetMangas() ([]*domain.Manga, error) {
	mangas, err := s.mangaRepo.List()
	if err != nil {
		return nil, err
	}

	// Sanitize all mangas
	sanitizedMangas := make([]*domain.Manga, len(mangas))
	for i, manga := range mangas {
		sanitizedMangas[i] = manga.Sanitize()
	}

	return sanitizedMangas, nil
}

// GetMangasByUser retrieves mangas by user ID
func (s *mangaService) GetMangasByUser(userID uint) ([]*domain.Manga, error) {
	mangas, err := s.mangaRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Sanitize all mangas
	sanitizedMangas := make([]*domain.Manga, len(mangas))
	for i, manga := range mangas {
		sanitizedMangas[i] = manga.Sanitize()
	}

	return sanitizedMangas, nil
}

// UpdateManga updates an existing manga
func (s *mangaService) UpdateManga(id uint, req *domain.UpdateMangaRequest, userID uint) (*domain.Manga, error) {
	// Get existing manga
	manga, err := s.mangaRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check ownership (user can only update their own manga)
	if manga.UserCreated != userID {
		return nil, errors.New("access denied: you can only update your own manga")
	}

	// Update manga fields
	manga.Name = req.Name
	manga.Price = req.Price
	manga.IsActive = req.IsActive

	if err := s.mangaRepo.Update(manga); err != nil {
		return nil, err
	}

	return manga.Sanitize(), nil
}

// DeleteManga deletes a manga by ID
func (s *mangaService) DeleteManga(id uint, userID uint) error {
	// Get existing manga
	manga, err := s.mangaRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Check ownership (user can only delete their own manga)
	if manga.UserCreated != userID {
		return errors.New("access denied: you can only delete your own manga")
	}

	return s.mangaRepo.Delete(id)
}

// GetActiveMangas retrieves all active mangas
func (s *mangaService) GetActiveMangas() ([]*domain.Manga, error) {
	mangas, err := s.mangaRepo.GetActiveMangas()
	if err != nil {
		return nil, err
	}

	// Sanitize all mangas
	sanitizedMangas := make([]*domain.Manga, len(mangas))
	for i, manga := range mangas {
		sanitizedMangas[i] = manga.Sanitize()
	}

	return sanitizedMangas, nil
}

// GetMangasByPriceRange retrieves mangas within price range
func (s *mangaService) GetMangasByPriceRange(min, max float64) ([]*domain.Manga, error) {
	mangas, err := s.mangaRepo.GetMangasByPriceRange(min, max)
	if err != nil {
		return nil, err
	}

	// Sanitize all mangas
	sanitizedMangas := make([]*domain.Manga, len(mangas))
	for i, manga := range mangas {
		sanitizedMangas[i] = manga.Sanitize()
	}

	return sanitizedMangas, nil
}

// GetMangasPaginated retrieves paginated mangas
func (s *mangaService) GetMangasPaginated(pagination *domain.PaginationRequest) (*domain.PaginatedResult[*domain.Manga], error) {
	mangas, total, err := s.mangaRepo.ListPaginated(pagination)
	if err != nil {
		return nil, err
	}

	// Sanitize all mangas
	sanitizedMangas := make([]*domain.Manga, len(mangas))
	for i, manga := range mangas {
		sanitizedMangas[i] = manga.Sanitize()
	}

	// Create pagination metadata
	paginationMeta := domain.NewPaginationResponse(pagination.Page, pagination.PageSize, total)

	return &domain.PaginatedResult[*domain.Manga]{
		Data:       sanitizedMangas,
		Pagination: paginationMeta,
	}, nil
}

// GetActiveMangasPaginated retrieves paginated active mangas
func (s *mangaService) GetActiveMangasPaginated(pagination *domain.PaginationRequest) (*domain.PaginatedResult[*domain.Manga], error) {
	mangas, total, err := s.mangaRepo.GetActiveMangasPaginated(pagination)
	if err != nil {
		return nil, err
	}

	// Sanitize all mangas
	sanitizedMangas := make([]*domain.Manga, len(mangas))
	for i, manga := range mangas {
		sanitizedMangas[i] = manga.Sanitize()
	}

	// Create pagination metadata
	paginationMeta := domain.NewPaginationResponse(pagination.Page, pagination.PageSize, total)

	return &domain.PaginatedResult[*domain.Manga]{
		Data:       sanitizedMangas,
		Pagination: paginationMeta,
	}, nil
}

// GetMangasByUserPaginated retrieves paginated mangas by user ID
func (s *mangaService) GetMangasByUserPaginated(userID uint, pagination *domain.PaginationRequest) (*domain.PaginatedResult[*domain.Manga], error) {
	mangas, total, err := s.mangaRepo.GetMangasByUserIDPaginated(userID, pagination)
	if err != nil {
		return nil, err
	}

	// Sanitize all mangas
	sanitizedMangas := make([]*domain.Manga, len(mangas))
	for i, manga := range mangas {
		sanitizedMangas[i] = manga.Sanitize()
	}

	// Create pagination metadata
	paginationMeta := domain.NewPaginationResponse(pagination.Page, pagination.PageSize, total)

	return &domain.PaginatedResult[*domain.Manga]{
		Data:       sanitizedMangas,
		Pagination: paginationMeta,
	}, nil
}

// GetMangasByPriceRangePaginated retrieves paginated mangas within price range
func (s *mangaService) GetMangasByPriceRangePaginated(min, max float64, pagination *domain.PaginationRequest) (*domain.PaginatedResult[*domain.Manga], error) {
	mangas, total, err := s.mangaRepo.GetMangasByPriceRangePaginated(min, max, pagination)
	if err != nil {
		return nil, err
	}

	// Sanitize all mangas
	sanitizedMangas := make([]*domain.Manga, len(mangas))
	for i, manga := range mangas {
		sanitizedMangas[i] = manga.Sanitize()
	}

	// Create pagination metadata
	paginationMeta := domain.NewPaginationResponse(pagination.Page, pagination.PageSize, total)

	return &domain.PaginatedResult[*domain.Manga]{
		Data:       sanitizedMangas,
		Pagination: paginationMeta,
	}, nil
}
