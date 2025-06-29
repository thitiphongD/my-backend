# 🏗️ My Backend - Project Overview

> **Go Fiber backend with Clean Architecture pattern** - Production-ready API server with proper dependency injection

## 🚀 **Quick Start**

```bash
# Clone & Setup
git clone <repository>
cd my-backend

# Environment Setup
cp .env.example .env  # Configure your environment variables

# Build & Run
go mod tidy                    # Download dependencies
go build -o bin/server ./cmd/server  # Build to bin/ directory
./bin/server                   # Run server

# Server starts on http://localhost:8080
# Health check: curl http://localhost:8080/
```

## 📁 **Project Structure (Clean Architecture)**

```
my-backend/
├── cmd/server/main.go           # 🚀 Application entry point & DI container
├── bin/                         # 📦 Build artifacts
│   ├── .gitkeep                 # Keeps directory in git
│   └── server                   # Binary (ignored by git)
├── internal/                    # 🏠 Private application code
│   ├── core/                    # 🎯 Business Logic Layer
│   │   ├── domain/              # Entities, DTOs & Value Objects
│   │   │   ├── user.go          # User entity with business rules
│   │   │   ├── manga.go         # Manga entity
│   │   │   ├── auth_dto.go      # Authentication DTOs
│   │   │   ├── manga_dto.go     # Manga DTOs
│   │   │   └── pagination.go    # Pagination domain objects
│   │   ├── ports/               # 🔌 Interfaces (Contracts)
│   │   │   ├── user_repository.go   # User data contracts
│   │   │   ├── manga_repository.go  # Manga data contracts
│   │   │   ├── auth_service.go      # Auth business contracts
│   │   │   └── manga_service.go     # Manga business contracts
│   │   └── services/            # 🔧 Business Logic Implementation
│   │       ├── user_service.go  # User business logic
│   │       ├── auth_service.go  # Authentication logic
│   │       └── manga_service.go # Manga business logic
│   ├── adapters/                # 🔌 Infrastructure Layer
│   │   ├── database/            # 🗄️ Database Infrastructure
│   │   │   ├── connection.go    # Database connection setup
│   │   │   └── repositories/    # Repository implementations
│   │   │       ├── user_repository.go   # User data access
│   │   │       └── manga_repository.go  # Manga data access
│   │   └── http/                # 🌐 HTTP Infrastructure
│   │       ├── handlers/        # HTTP request handlers
│   │       │   ├── user_handler.go  # User HTTP handlers
│   │       │   ├── auth_handler.go  # Auth HTTP handlers
│   │       │   └── manga_handler.go # Manga HTTP handlers
│   │       ├── middleware/      # HTTP middleware
│   │       │   └── auth_middleware.go # JWT authentication
│   │       └── routes/          # Route configuration
│   │           └── routes.go    # All route definitions
│   ├── config/                  # ⚙️ Configuration management
│   │   └── config.go            # Environment config loading
│   └── utils/                   # 🔧 Shared utilities
│       ├── jwt.go               # JWT token utilities
│       └── password.go          # Password hashing utilities
├── pkg/                         # 📦 Public packages (reusable)
│   ├── response/                # HTTP response helpers
│   └── validator/               # Input validation utilities
├── .env                         # 🔒 Environment variables (ignored)
├── .gitignore                   # 📋 Git ignore patterns
└── go.mod                       # 📦 Go module definition
```

## 🌐 **Available APIs**

### **Authentication**
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login  
- `GET /api/v1/auth/me` - Get current user (protected)

### **User Management**
- `GET /api/v1/users` - List all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create user (protected)
- `PUT /api/v1/users/:id` - Update user (protected)
- `DELETE /api/v1/users/:id` - Delete user (protected)

### **Manga Management**
- `GET /api/v1/mangas` - List all mangas
- `GET /api/v1/mangas/:id` - Get manga by ID
- `POST /api/v1/mangas` - Create manga (protected)
- `PUT /api/v1/mangas/:id` - Update manga (protected)
- `DELETE /api/v1/mangas/:id` - Delete manga (protected)

### **Pagination Support**
- `GET /api/v1/mangas/paginated` - Paginated manga list
- `GET /api/v1/mangas/active/paginated` - Paginated active mangas
- `GET /api/v1/mangas/user/:id/paginated` - Paginated user mangas

## ⚙️ **Configuration**

### **Environment Variables**
```bash
# Server Configuration
PORT=8080



# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key

# Optional: Individual DB settings (if not using connection string)
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=mydb
DB_SSL_MODE=require
DB_CHANNEL_BINDING=require
```

### **Configuration Loading**
- Environment variables loaded with fallback defaults
- `.env` file support for development
- `config.LoadConfig()` centralizes configuration management
- JWT secret validation with warnings for insecure defaults

## 🎯 **Key Features**

✅ **Clean Architecture** - Separation of concerns, testable code  
✅ **Dependency Injection** - Proper DI container in main.go  
✅ **JWT Authentication** - Secure user authentication  
✅ **CRUD Operations** - Full Create, Read, Update, Delete support  
✅ **Pagination System** - Efficient data loading with metadata  
✅ **Input Validation** - Request validation with error handling  
✅ **Database Integration** - PostgreSQL with GORM  
✅ **Middleware Support** - Authentication, CORS, logging  
✅ **Error Handling** - Consistent error responses  
✅ **Build Management** - Proper binary management with .gitignore

## 🏗️ **Clean Architecture Benefits**

### **🎯 Business Logic Protection**
- **Domain entities** contain business rules and validation
- **Services** implement use cases without infrastructure dependencies
- **Ports (interfaces)** define contracts between layers

### **🔧 Infrastructure Flexibility**
- **Database** can be swapped (PostgreSQL → MySQL → MongoDB)
- **HTTP framework** can be changed (Fiber → Gin → Chi)
- **Authentication** method can be modified (JWT → OAuth → Session)

### **🧪 Testing Excellence**
- **Unit tests** for business logic (services, entities)
- **Integration tests** for adapters (repositories, handlers)
- **Mocking** made easy with interface-based design

### **👥 Team Collaboration**
- **Clear boundaries** between team responsibilities
- **Parallel development** - frontend, backend, database teams
- **Code reviews** focused on specific layers  

## 📚 **Documentation**

| **Guide** | **Purpose** | **For Who** |
|-----------|-------------|-------------|
| [📄 Pagination Guide](PAGINATION_GUIDE.md) | Pagination usage | API users |
| [🛠️ How to Add New API](HOW_TO_ADD_NEW_API.md) | Step-by-step implementation | Developers |
| [🎓 Clean Architecture Flow](CLEAN_ARCHITECTURE_FLOW.md) | Architecture deep dive | Architects |

## 🧪 **Response Format**

### **Success Response**
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

### **Error Response**
```json
{
  "success": false,
  "error": "Error message"
}
```

### **Paginated Response**
```json
{
  "success": true,
  "message": "Paginated data retrieved successfully",
  "data": {
    "data": [ ... ],
    "pagination": {
      "current_page": 1,
      "page_size": 10,
      "total_items": 25,
      "total_pages": 3,
      "has_next_page": true,
      "has_prev_page": false
    }
  }
}
```

## 🛠️ **Development & Deployment**

### **Local Development**
```bash
# Setup
go mod tidy
cp .env.example .env

# Development build
go build -o bin/server ./cmd/server

# Run with auto-reload (using air - optional)
air

# Testing
go test ./... -v

# Check architecture compliance
go mod graph | grep internal  # Should not import external deps
```

### **Build & Deployment**
```bash
# Production build
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server

# Docker build (using provided Dockerfile)
docker build -t my-backend .
docker run -p 8080:8080 --env-file .env my-backend

# Binary deployment
./bin/server  # Ensure environment variables are set
```

### **File Management**
- **Binary files** are built to `bin/` directory
- **`.gitignore`** properly excludes build artifacts
- **`bin/.gitkeep`** maintains directory structure in git
- **Environment files** (`.env`) are never committed

## 📝 **Next Steps**

1. **For API Usage**: Start with [Pagination Guide](PAGINATION_GUIDE.md)
2. **For Development**: Follow [How to Add New API](HOW_TO_ADD_NEW_API.md)  
3. **For Architecture**: Study [Clean Architecture Flow](CLEAN_ARCHITECTURE_FLOW.md)

---

🚀 **Ready to use!** This backend provides a solid foundation for building scalable REST APIs with clean, maintainable code. 