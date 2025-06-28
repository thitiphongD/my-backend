# 🏗️ Clean Architecture - Project Structure

โปรเจ็กต์นี้ได้รับการ refactor ให้ใช้ **Clean Architecture Pattern** ตามแนวทาง **Google Developer Best Practices**

## 📁 โครงสร้างใหม่

```
my-backend/
├── cmd/                              # Entry points
│   └── server/
│       └── main.go                   # Application entry point
├── internal/                         # Private application code
│   ├── adapters/                     # Infrastructure Layer
│   │   ├── database/
│   │   │   └── repositories/         # Data access implementations
│   │   └── http/                     # HTTP infrastructure
│   │       ├── handlers/             # HTTP handlers (controllers)
│   │       ├── middleware/           # HTTP middlewares
│   │       └── routes/               # Route definitions
│   ├── config/                       # Configuration management
│   └── core/                         # Business Logic Layer
│       ├── domain/                   # Entities & DTOs
│       ├── ports/                    # Interfaces (contracts)
│       └── services/                 # Business logic implementation
├── pkg/                              # Public/shared packages
│   ├── response/                     # HTTP response helpers
│   └── validator/                    # Input validation
└── internal/ (legacy)                # เก่า - จะลบออกได้
    ├── auth/
    ├── database/
    ├── middleware/
    ├── models/
    └── utils/
```

## 🎯 Architecture Layers

### 1. 🌐 **Presentation Layer** (`/internal/adapters/http/`)
- **Handlers**: จัดการ HTTP requests/responses
- **Middleware**: Authentication, CORS, logging
- **Routes**: การจัดกลุ่มและกำหนด routes

### 2. 📋 **Application Layer** (`/internal/core/services/`)
- **Services**: Business logic และ use cases
- **Dependency Injection**: Services ขึ้นกับ interfaces ไม่ใช่ concrete implementations

### 3. 🏢 **Domain Layer** (`/internal/core/domain/`)
- **Entities**: Core business objects (User)
- **DTOs**: Data Transfer Objects สำหรับ API
- **Business Rules**: Validation และ business logic ใน entities

### 4. 🔌 **Ports** (`/internal/core/ports/`)
- **Interfaces**: Contracts ระหว่าง layers
- **Repository Interfaces**: กำหนด data access contracts
- **Service Interfaces**: กำหนด business logic contracts

### 5. 🔧 **Infrastructure Layer** (`/internal/adapters/`)
- **Repository Implementations**: การเชื่อมต่อ database จริง
- **Database**: Connection และ migration
- **External Services**: Third-party integrations

## ✨ ข้อดีของโครงสร้างใหม่

### 🎯 **Separation of Concerns**
- แต่ละ layer มีหน้าที่ชัดเจน
- Business logic แยกออกจาก infrastructure
- ง่ายต่อการ maintain และ debug

### 🧪 **Testability**
- สามารถ mock dependencies ได้ง่าย
- Unit testing แต่ละ layer แยกกันได้
- Integration testing ที่มีประสิทธิภาพ

### 🔄 **Dependency Inversion**
- High-level modules ไม่ขึ้นกับ low-level modules
- ทั้งคู่ขึ้นกับ abstractions (interfaces)
- ง่ายต่อการเปลี่ยน implementation

### 📈 **Scalability**
- เพิ่ม features ใหม่โดยไม่กระทบของเก่า
- แยก modules ได้อย่างชัดเจน
- Support microservices architecture

## 🚀 การใช้งาน

### Build & Run
```bash
# Build
go build -o bin/server ./cmd/server

# Run
./bin/server
# หรือ
go run ./cmd/server
```

### API Endpoints

#### **Authentication**
- `POST /api/v1/auth/register` - สมัครสมาชิก
- `POST /api/v1/auth/login` - เข้าสู่ระบบ
- `GET /api/v1/auth/me` - ข้อมูลผู้ใช้ปัจจุบัน (ต้อง login)

#### **User Management**
- `GET /api/v1/users` - รายการผู้ใช้ทั้งหมด
- `GET /api/v1/users/:id` - ข้อมูลผู้ใช้ตาม ID
- `POST /api/v1/users` - สร้างผู้ใช้ใหม่ (ต้อง login)
- `PUT /api/v1/users/:id` - แก้ไขข้อมูลผู้ใช้ (ต้อง login)
- `DELETE /api/v1/users/:id` - ลบผู้ใช้ (ต้อง login)

### Response Format
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... },
  "error": null
}
```

## 🔧 Configuration

### Environment Variables
```bash
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=mydb
JWT_SECRET=your-secret-key
```

## 📋 TODO: การปรับปรุงต่อไป

1. **ลบโค้ดเก่า**: ลบ `/internal/auth/`, `/internal/models/` ที่ไม่ใช้แล้ว
2. **เพิ่ม Unit Tests**: สร้าง tests สำหรับแต่ละ layer
3. **เพิ่ม Swagger Documentation**: API documentation
4. **เพิ่ม Logging**: Structured logging ด้วย logrus หรือ zap
5. **เพิ่ม Database Migrations**: Proper migration system
6. **เพิ่ม Docker Support**: Containerization
7. **เพิ่ม CI/CD Pipeline**: Automated testing และ deployment

## 🎓 Patterns ที่ใช้

- **Clean Architecture** (Uncle Bob)
- **Repository Pattern** (Data Access Layer)
- **Dependency Injection** (IoC Container)
- **Factory Pattern** (Service creation)
- **Middleware Pattern** (Cross-cutting concerns)
- **DTO Pattern** (Data Transfer Objects) 