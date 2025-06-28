# 🚀 My Fiber App – Web API built with Golang & Fiber

This is a web application built with [Golang](https://golang.org) using the [Fiber](https://github.com/gofiber/fiber) web framework. Fiber is a lightweight and high-performance HTTP framework that helps you build web applications and REST APIs efficiently.

# My Backend API

Go Fiber backend with PostgreSQL database, JWT authentication, and CORS support.

## Features

- ✅ RESTful API with Go Fiber
- ✅ PostgreSQL database with GORM
- ✅ JWT Authentication middleware
- ✅ CORS middleware
- ✅ Password hashing with bcrypt
- ✅ Request logging
- ✅ Environment configuration

## Setup

### 1. Environment Variables

สร้างไฟล์ `.env` จาก `.env.example`:

```bash
cp .env.example .env
```

แก้ไขค่าในไฟล์ `.env`:

```env
PORT=8080
DB_CONNECTION_STRING=postgresql://username:password@localhost:5432/database_name?sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Application

```bash
go run main.go
```

## API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Welcome message |
| GET | `/health` | Database health check |
| POST | `/auth/register` | User registration |
| POST | `/auth/login` | User login |
| GET | `/users` | Get all users |
| GET | `/users/:id` | Get user by ID |

### Protected Endpoints (require JWT token)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/auth/me` | Get current user profile |
| POST | `/users` | Create new user (admin) |

## Authentication

### Register

```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Using Protected Endpoints

ใช้ JWT token ที่ได้รับจาก login/register ใน Authorization header:

```bash
curl -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  http://localhost:8080/auth/me
```

## CORS Configuration

CORS ถูกตั้งค่าให้รองรับ:
- React dev server (`http://localhost:3000`)
- Vite dev server (`http://localhost:5173`)
- All origins (`*`) for development

## Database Schema

### User Model

```go
type User struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"-"` // Hidden from JSON response
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}


Build Command: go build -o app cmd/server/main.go
Start Command: ./app
```
