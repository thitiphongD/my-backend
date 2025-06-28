# 🛠️ **How to Add New API - Step by Step Guide**

> **Practical implementation guide** for adding new APIs following Clean Architecture pattern

## 🎯 **Prerequisites**

- Go 1.19+ installed
- PostgreSQL database running
- Basic understanding of Go and REST APIs
- Familiarity with project structure (see [ARCHITECTURE.md](ARCHITECTURE.md))

---

## 📋 **Implementation Checklist**

ตัวอย่าง: เพิ่ม **Book CRUD API**

```bash
# Fields: title (string), author (string), price (float), available (bool)
# Operations: Create, Read, Update, Delete
```

### **✅ Step-by-Step Checklist**

| Step | File | Status |
|------|------|--------|
| 1 | `internal/core/domain/book.go` | ⬜ Domain Entity |
| 2 | `internal/core/domain/book_dto.go` | ⬜ DTOs |
| 3 | `internal/core/ports/book_repository.go` | ⬜ Repository Interface |
| 4 | `internal/core/ports/book_service.go` | ⬜ Service Interface |
| 5 | `internal/adapters/database/repositories/book_repository.go` | ⬜ Repository Implementation |
| 6 | `internal/core/services/book_service.go` | ⬜ Service Implementation |
| 7 | `internal/adapters/http/handlers/book_handler.go` | ⬜ HTTP Handler |
| 8 | `internal/adapters/http/routes/routes.go` | ⬜ Routes Setup |
| 9 | `cmd/server/main.go` | ⬜ Dependency Injection |

---

## 🚀 **Step-by-Step Implementation**

### **Step 1: Domain Entity** 📦
📁 `internal/core/domain/book.go`

```go
package domain

import (
	"strings"
	"time"
	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Title     string         `json:"title" gorm:"not null"`
	Author    string         `json:"author" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	Available bool           `json:"available" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// IsValid validates book data
func (b *Book) IsValid() bool {
	return b.Title != "" && b.Author != "" && b.Price >= 0
}

// Sanitize cleans book data
func (b *Book) Sanitize() *Book {
	sanitized := *b
	sanitized.Title = strings.TrimSpace(sanitized.Title)
	sanitized.Author = strings.TrimSpace(sanitized.Author)
	return &sanitized
}
```

### **Step 2: DTOs** 📤
📁 `internal/core/domain/book_dto.go`

```go
package domain

type CreateBookRequest struct {
	Title     string  `json:"title" validate:"required,min=1,max=200"`
	Author    string  `json:"author" validate:"required,min=1,max=100"`
	Price     float64 `json:"price" validate:"required,min=0"`
	Available bool    `json:"available"`
}

type UpdateBookRequest struct {
	Title     string  `json:"title" validate:"required,min=1,max=200"`
	Author    string  `json:"author" validate:"required,min=1,max=100"`
	Price     float64 `json:"price" validate:"required,min=0"`
	Available bool    `json:"available"`
}

type BookResponse struct {
	ID        uint    `json:"id"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
```

### **Step 3: Repository Interface** 🔌
📁 `internal/core/ports/book_repository.go`

```go
package ports

import "github.com/yourproject/internal/core/domain"

type BookRepository interface {
	Create(book *domain.Book) error
	GetByID(id uint) (*domain.Book, error)
	List() ([]*domain.Book, error)
	Update(book *domain.Book) error
	Delete(id uint) error
	
	// Business specific queries
	GetAvailableBooks() ([]*domain.Book, error)
	GetBooksByAuthor(author string) ([]*domain.Book, error)
	SearchBooks(query string) ([]*domain.Book, error)
}
```

### **Step 4: Service Interface** 🔌
📁 `internal/core/ports/book_service.go`

```go
package ports

import "github.com/yourproject/internal/core/domain"

type BookService interface {
	CreateBook(req *domain.CreateBookRequest) (*domain.Book, error)
	GetBookByID(id uint) (*domain.Book, error)
	GetBooks() ([]*domain.Book, error)
	UpdateBook(id uint, req *domain.UpdateBookRequest) (*domain.Book, error)
	DeleteBook(id uint) error
	
	// Business operations
	GetAvailableBooks() ([]*domain.Book, error)
	GetBooksByAuthor(author string) ([]*domain.Book, error)
	SearchBooks(query string) ([]*domain.Book, error)
}
```

### **Step 5: Repository Implementation** 🗄️
📁 `internal/adapters/database/repositories/book_repository.go`

```go
package repositories

import (
	"errors"
	"github.com/yourproject/internal/core/domain"
	"github.com/yourproject/internal/core/ports"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) ports.BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(book *domain.Book) error {
	if err := r.db.Create(book).Error; err != nil {
		return errors.New("failed to create book")
	}
	return nil
}

func (r *bookRepository) GetByID(id uint) (*domain.Book, error) {
	var book domain.Book
	if err := r.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, errors.New("failed to get book")
	}
	return &book, nil
}

func (r *bookRepository) List() ([]*domain.Book, error) {
	var books []*domain.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, errors.New("failed to get books")
	}
	return books, nil
}

func (r *bookRepository) Update(book *domain.Book) error {
	if err := r.db.Save(book).Error; err != nil {
		return errors.New("failed to update book")
	}
	return nil
}

func (r *bookRepository) Delete(id uint) error {
	if err := r.db.Delete(&domain.Book{}, id).Error; err != nil {
		return errors.New("failed to delete book")
	}
	return nil
}

func (r *bookRepository) GetAvailableBooks() ([]*domain.Book, error) {
	var books []*domain.Book
	if err := r.db.Where("available = ?", true).Find(&books).Error; err != nil {
		return nil, errors.New("failed to get available books")
	}
	return books, nil
}

func (r *bookRepository) GetBooksByAuthor(author string) ([]*domain.Book, error) {
	var books []*domain.Book
	if err := r.db.Where("author ILIKE ?", "%"+author+"%").Find(&books).Error; err != nil {
		return nil, errors.New("failed to get books by author")
	}
	return books, nil
}

func (r *bookRepository) SearchBooks(query string) ([]*domain.Book, error) {
	var books []*domain.Book
	searchPattern := "%" + query + "%"
	if err := r.db.Where("title ILIKE ? OR author ILIKE ?", searchPattern, searchPattern).Find(&books).Error; err != nil {
		return nil, errors.New("failed to search books")
	}
	return books, nil
}
```

### **Step 6: Service Implementation** 🔧
📁 `internal/core/services/book_service.go`

```go
package services

import (
	"errors"
	"github.com/yourproject/internal/core/domain"
	"github.com/yourproject/internal/core/ports"
)

type bookService struct {
	bookRepo ports.BookRepository
}

func NewBookService(bookRepo ports.BookRepository) ports.BookService {
	return &bookService{bookRepo: bookRepo}
}

func (s *bookService) CreateBook(req *domain.CreateBookRequest) (*domain.Book, error) {
	book := &domain.Book{
		Title:     req.Title,
		Author:    req.Author,
		Price:     req.Price,
		Available: req.Available,
	}

	// Validate business rules
	if !book.IsValid() {
		return nil, errors.New("invalid book data")
	}

	// Sanitize data
	book = book.Sanitize()

	if err := s.bookRepo.Create(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookService) GetBookByID(id uint) (*domain.Book, error) {
	book, err := s.bookRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return book.Sanitize(), nil
}

func (s *bookService) GetBooks() ([]*domain.Book, error) {
	books, err := s.bookRepo.List()
	if err != nil {
		return nil, err
	}

	// Sanitize all books
	sanitizedBooks := make([]*domain.Book, len(books))
	for i, book := range books {
		sanitizedBooks[i] = book.Sanitize()
	}

	return sanitizedBooks, nil
}

func (s *bookService) UpdateBook(id uint, req *domain.UpdateBookRequest) (*domain.Book, error) {
	book, err := s.bookRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	book.Title = req.Title
	book.Author = req.Author
	book.Price = req.Price
	book.Available = req.Available

	// Validate business rules
	if !book.IsValid() {
		return nil, errors.New("invalid book data")
	}

	// Sanitize data
	book = book.Sanitize()

	if err := s.bookRepo.Update(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookService) DeleteBook(id uint) error {
	// Check if book exists
	_, err := s.bookRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.bookRepo.Delete(id)
}

func (s *bookService) GetAvailableBooks() ([]*domain.Book, error) {
	books, err := s.bookRepo.GetAvailableBooks()
	if err != nil {
		return nil, err
	}

	// Sanitize all books
	sanitizedBooks := make([]*domain.Book, len(books))
	for i, book := range books {
		sanitizedBooks[i] = book.Sanitize()
	}

	return sanitizedBooks, nil
}

func (s *bookService) GetBooksByAuthor(author string) ([]*domain.Book, error) {
	if author == "" {
		return nil, errors.New("author name is required")
	}

	books, err := s.bookRepo.GetBooksByAuthor(author)
	if err != nil {
		return nil, err
	}

	// Sanitize all books
	sanitizedBooks := make([]*domain.Book, len(books))
	for i, book := range books {
		sanitizedBooks[i] = book.Sanitize()
	}

	return sanitizedBooks, nil
}

func (s *bookService) SearchBooks(query string) ([]*domain.Book, error) {
	if query == "" {
		return nil, errors.New("search query is required")
	}

	books, err := s.bookRepo.SearchBooks(query)
	if err != nil {
		return nil, err
	}

	// Sanitize all books
	sanitizedBooks := make([]*domain.Book, len(books))
	for i, book := range books {
		sanitizedBooks[i] = book.Sanitize()
	}

	return sanitizedBooks, nil
}
```

### **Step 7: HTTP Handler** 🌐
📁 `internal/adapters/http/handlers/book_handler.go`

```go
package handlers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/yourproject/internal/core/domain"
	"github.com/yourproject/internal/core/ports"
	"github.com/yourproject/pkg/response"
	"github.com/yourproject/pkg/validator"
)

type BookHandler struct {
	bookService ports.BookService
}

func NewBookHandler(bookService ports.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

// CreateBook godoc
// @Summary Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body domain.CreateBookRequest true "Book data"
// @Success 201 {object} response.Response{data=domain.Book}
// @Router /api/v1/books [post]
func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var req domain.CreateBookRequest
	if err := validator.ParseAndValidate(c, &req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid request data")
	}

	book, err := h.bookService.CreateBook(&req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Failed to create book")
	}

	return response.Created(c, book, "Book created successfully")
}

// GetBooks godoc
// @Summary Get all books
// @Tags books
// @Produce json
// @Success 200 {object} response.Response{data=[]domain.Book}
// @Router /api/v1/books [get]
func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.bookService.GetBooks()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get books")
	}

	return response.Success(c, books, "Books retrieved successfully")
}

// GetBook godoc
// @Summary Get book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.Response{data=domain.Book}
// @Router /api/v1/books/{id} [get]
func (h *BookHandler) GetBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid book ID")
	}

	book, err := h.bookService.GetBookByID(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err, "Book not found")
	}

	return response.Success(c, book, "Book retrieved successfully")
}

// UpdateBook godoc
// @Summary Update book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body domain.UpdateBookRequest true "Updated book data"
// @Success 200 {object} response.Response{data=domain.Book}
// @Router /api/v1/books/{id} [put]
func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid book ID")
	}

	var req domain.UpdateBookRequest
	if err := validator.ParseAndValidate(c, &req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid request data")
	}

	book, err := h.bookService.UpdateBook(uint(id), &req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Failed to update book")
	}

	return response.Success(c, book, "Book updated successfully")
}

// DeleteBook godoc
// @Summary Delete book
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.Response
// @Router /api/v1/books/{id} [delete]
func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid book ID")
	}

	if err := h.bookService.DeleteBook(uint(id)); err != nil {
		return response.Error(c, fiber.StatusNotFound, err, "Failed to delete book")
	}

	return response.Success(c, nil, "Book deleted successfully")
}

// GetAvailableBooks godoc
// @Summary Get available books
// @Tags books
// @Produce json
// @Success 200 {object} response.Response{data=[]domain.Book}
// @Router /api/v1/books/available [get]
func (h *BookHandler) GetAvailableBooks(c *fiber.Ctx) error {
	books, err := h.bookService.GetAvailableBooks()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get available books")
	}

	return response.Success(c, books, "Available books retrieved successfully")
}

// GetBooksByAuthor godoc
// @Summary Get books by author
// @Tags books
// @Produce json
// @Param author query string true "Author name"
// @Success 200 {object} response.Response{data=[]domain.Book}
// @Router /api/v1/books/author [get]
func (h *BookHandler) GetBooksByAuthor(c *fiber.Ctx) error {
	author := c.Query("author")
	if author == "" {
		return response.Error(c, fiber.StatusBadRequest, nil, "Author parameter is required")
	}

	books, err := h.bookService.GetBooksByAuthor(author)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get books by author")
	}

	return response.Success(c, books, "Books by author retrieved successfully")
}

// SearchBooks godoc
// @Summary Search books
// @Tags books
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} response.Response{data=[]domain.Book}
// @Router /api/v1/books/search [get]
func (h *BookHandler) SearchBooks(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return response.Error(c, fiber.StatusBadRequest, nil, "Search query parameter is required")
	}

	books, err := h.bookService.SearchBooks(query)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to search books")
	}

	return response.Success(c, books, "Books search completed successfully")
}
```

### **Step 8: Routes Setup** 🛣️
แก้ไข `internal/adapters/http/routes/routes.go`

```go
func SetupRoutes(app *fiber.App, authService ports.AuthService, userService ports.UserService, mangaService ports.MangaService, bookService ports.BookService) {
	// ... existing code ...

	// Initialize book handler
	bookHandler := handlers.NewBookHandler(bookService)

	// Book routes
	books := v1.Group("/books")
	books.Get("/", bookHandler.GetBooks)                    // Public: Get all books
	books.Get("/available", bookHandler.GetAvailableBooks)  // Public: Get available books
	books.Get("/author", bookHandler.GetBooksByAuthor)      // Public: Get books by author
	books.Get("/search", bookHandler.SearchBooks)           // Public: Search books
	books.Get("/:id", bookHandler.GetBook)                  // Public: Get book by ID
	books.Post("/", bookHandler.CreateBook)                 // Public: Create book
	books.Put("/:id", bookHandler.UpdateBook)               // Public: Update book
	books.Delete("/:id", bookHandler.DeleteBook)            // Public: Delete book
}
```

### **Step 9: Dependency Injection** ⚙️
แก้ไข `cmd/server/main.go`

```go
func main() {
	// ... existing code ...

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	mangaRepo := repositories.NewMangaRepository(db)
	bookRepo := repositories.NewBookRepository(db) // ➕ Add this

	// Initialize services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	mangaService := services.NewMangaService(mangaRepo)
	bookService := services.NewBookService(bookRepo) // ➕ Add this

	// Setup routes
	routes.SetupRoutes(app, authService, userService, mangaService, bookService) // ➕ Add bookService

	// Auto migrate
	if err := db.AutoMigrate(&domain.User{}, &domain.Manga{}, &domain.Book{}); err != nil { // ➕ Add &domain.Book{}
		log.Fatal("Failed to migrate database: ", err)
	}

	// ... rest of code ...
}
```

---

## 🧪 **Testing Your New API**

### **1. Test Book Creation:**
```bash
curl -X POST "http://localhost:8080/api/v1/books" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "price": 45.99,
    "available": true
  }'
```

### **2. Test Get All Books:**
```bash
curl "http://localhost:8080/api/v1/books"
```

### **3. Test Search Books:**
```bash
curl "http://localhost:8080/api/v1/books/search?q=Clean"
```

### **4. Test Get Books by Author:**
```bash
curl "http://localhost:8080/api/v1/books/author?author=Robert"
```

### **5. Test Update Book:**
```bash
curl -X PUT "http://localhost:8080/api/v1/books/1" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Clean Code - Updated",
    "author": "Robert C. Martin",
    "price": 49.99,
    "available": true
  }'
```

### **6. Test Delete Book:**
```bash
curl -X DELETE "http://localhost:8080/api/v1/books/1"
```

---

## ✅ **Completion Checklist**

หลังจากทำทุก step แล้ว ตรวจสอบ:

- [ ] **Compile Successfully**: `go build ./cmd/server`
- [ ] **Database Migration**: เห็น `books` table ใน DB
- [ ] **API Responses**: ทุก endpoint ตอบ status 200/201
- [ ] **Data Validation**: ทดสอบส่งข้อมูลผิด format
- [ ] **Error Handling**: ทดสอบ edge cases (book not found, etc.)
- [ ] **Search Functionality**: ทดสอบค้นหาหนังสือ
- [ ] **Business Logic**: ทดสอบ validation rules

---

## 🎉 **Success!**

ยินดีด้วย! คุณได้เพิ่ม Book API สำเร็จแล้ว 

**📚 Next Learning:**
- ลองเพิ่ม **Pagination** ให้ Book API (ดู [PAGINATION_GUIDE.md](PAGINATION_GUIDE.md))
- ศึกษา **Clean Architecture** ลึกขึ้น (ดู [CLEAN_ARCHITECTURE_FLOW.md](CLEAN_ARCHITECTURE_FLOW.md))
- เพิ่ม **Unit Tests** สำหรับ Book services และ handlers

**💡 Pro Tips:**
1. ใช้ `git commit` หลังจากทำแต่ละ step เสร็จ
2. ทดสอบ API ด้วย Postman หรือ cURL หลังจากแต่ละ step
3. ถ้า error ให้ดู logs ใน terminal และแก้ทีละตัว
4. อย่าลืม add validation tags ใน DTOs เพื่อป้องกัน invalid data

---

🚀 **Happy Coding!** Clean Architecture ทำให้การเพิ่ม API ใหม่เป็นเรื่องง่ายและมีโครงสร้างที่ชัดเจน!
