# 🏗️ My Backend - Project Overview

> **Go Fiber backend with Clean Architecture pattern** - Production-ready API server

## 🚀 **Quick Start**

```bash
# Clone & Setup
git clone <repository>
cd my-backend

# Build & Run
go build -o bin/server ./cmd/server
./bin/server

# Server starts on http://localhost:8080
```

## 📁 **Project Structure**

```
my-backend/
├── cmd/server/main.go           # 🚀 Application entry point
├── internal/core/               # 🏢 Business Logic
│   ├── domain/                  # Entities & DTOs
│   ├── ports/                   # Interfaces (contracts)
│   └── services/                # Business logic implementation
├── internal/adapters/           # 🔌 Infrastructure
│   ├── database/repositories/   # Data access
│   └── http/                    # HTTP handlers & routes
└── pkg/                         # 📦 Shared utilities
    ├── response/                # HTTP response helpers
    └── validator/               # Input validation
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

```bash
# Required Environment Variables
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=mydb
JWT_SECRET=your-secret-key
```

## 🎯 **Key Features**

✅ **Clean Architecture** - Separation of concerns, testable code  
✅ **JWT Authentication** - Secure user authentication  
✅ **CRUD Operations** - Full Create, Read, Update, Delete support  
✅ **Pagination System** - Efficient data loading  
✅ **Input Validation** - Request validation with error handling  
✅ **Database Integration** - PostgreSQL with GORM  
✅ **Middleware Support** - Authentication, CORS, logging  

## 📚 **Documentation**

| **Guide** | **Purpose** | **For Who** |
|-----------|-------------|-------------|
| [📄 Pagination Guide](PAGINATION_GUIDE.md) | Pagination usage | API users |
| [🛠️ How to Add New API](HOW_TO_ADD_NEW_API.md) | Step-by-step implementation | Developers |
| [🎓 Clean Architecture Flow](CLEAN_ARCHITECTURE_FLOW.md) | Architecture deep dive | Architects |

## 🧪 **Response Format**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

## 📝 **Next Steps**

1. **For API Usage**: Start with [Pagination Guide](PAGINATION_GUIDE.md)
2. **For Development**: Follow [How to Add New API](HOW_TO_ADD_NEW_API.md)  
3. **For Architecture**: Study [Clean Architecture Flow](CLEAN_ARCHITECTURE_FLOW.md)

---

🚀 **Ready to use!** This backend provides a solid foundation for building scalable REST APIs with clean, maintainable code. 