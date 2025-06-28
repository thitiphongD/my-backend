# 📄 **Pagination System Guide**

> **การใช้งานระบบ Pagination ใน Manga API** - Clean Architecture Pattern

## 🎯 **Overview**

ระบบ Pagination ได้รับการออกแบบตาม **Clean Architecture** โดยมีจุดประสงค์:
- **Performance**: ลดการโหลดข้อมูลจำนวนมากครั้งเดียว
- **User Experience**: แสดงข้อมูลเป็นหน้าๆ ง่ายต่อการใช้งาน  
- **Scalability**: รองรับข้อมูลขนาดใหญ่ในอนาคต
- **Consistency**: API response format เป็นมาตรฐานทุก endpoint

---

## 🏗️ **Pagination Architecture**

ระบบ Pagination ถูกออกแบบให้เข้ากับ Clean Architecture:

```
🔄 Request Flow:
HTTP Request → Handler → Service → Repository → Database
             ↓ Validate  ↓ Business  ↓ Query      ↓ Paginated Data
             Parameters   Logic       Execution    Results

📁 Implementation:
├── 📦 Domain: PaginationRequest, PaginationResponse, PaginatedResult[T]
├── 🔌 Ports: Repository & Service interfaces with pagination methods
├── 🗄️ Adapters: Database pagination queries (LIMIT, OFFSET)
└── 🌐 Handlers: HTTP parameter parsing & response formatting
```

---

## 📊 **Pagination Data Types**

### **PaginationRequest**
```go
type PaginationRequest struct {
    Page     int `query:"page" validate:"min=1"`
    PageSize int `query:"page_size" validate:"min=1,max=100"`
}
```

### **PaginationResponse** 
```go
type PaginationResponse struct {
    CurrentPage  int   `json:"current_page"`
    PageSize     int   `json:"page_size"`
    TotalItems   int64 `json:"total_items"`
    TotalPages   int   `json:"total_pages"`
    HasNextPage  bool  `json:"has_next_page"`
    HasPrevPage  bool  `json:"has_prev_page"`
    NextPage     *int  `json:"next_page,omitempty"`
    PreviousPage *int  `json:"previous_page,omitempty"`
}
```

### **PaginatedResult[T]**
```go
type PaginatedResult[T any] struct {
    Data       []T                 `json:"data"`
    Pagination *PaginationResponse `json:"pagination"`
}
```

---

## 🌐 **Available Pagination Endpoints**

| **Endpoint** | **Description** | **Parameters** |
|--------------|-----------------|----------------|
| `/api/v1/mangas/paginated` | All mangas (paginated) | `page`, `page_size` |
| `/api/v1/mangas/active/paginated` | Active mangas only | `page`, `page_size` |
| `/api/v1/mangas/user/:userID/paginated` | User's mangas | `page`, `page_size` |
| `/api/v1/mangas/price/paginated` | Price range mangas | `page`, `page_size`, `min`, `max` |

---

## 📋 **Query Parameters**

| **Parameter** | **Type** | **Default** | **Validation** | **Description** |
|---------------|----------|-------------|----------------|-----------------|
| `page` | `int` | `1` | `min=1` | Page number (1-based) |
| `page_size` | `int` | `10` | `min=1, max=100` | Items per page |
| `min` | `float64` | `0` | `>=0` | Minimum price (price range only) |
| `max` | `float64` | `999999` | `>=0` | Maximum price (price range only) |

---

## 🧪 **Usage Examples**

### **1. Basic Pagination** 📄

```bash
# Get first page with 5 items
curl "http://localhost:8080/api/v1/mangas/paginated?page=1&page_size=5"

# Response
{
  "success": true,
  "message": "Paginated mangas retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "name": "One Piece Updated",
        "price": 280,
        "is_active": true,
        "user_created": 4,
        "created_at": "2025-06-28T14:43:25.845583+07:00",
        "updated_at": "2025-06-28T14:44:25.959817+07:00"
      }
    ],
    "pagination": {
      "current_page": 1,
      "page_size": 5,
      "total_items": 10,
      "total_pages": 2,
      "has_next_page": true,
      "has_prev_page": false,
      "next_page": 2
    }
  }
}
```

### **2. Active Mangas Pagination** ✅

```bash
# Get active mangas only
curl "http://localhost:8080/api/v1/mangas/active/paginated?page=1&page_size=3"
```

### **3. User-Specific Pagination** 👤

```bash
# Get mangas by specific user
curl "http://localhost:8080/api/v1/mangas/user/4/paginated?page=1&page_size=10"
```

### **4. Price Range Pagination** 💰

```bash
# Get mangas within price range 100-500
curl "http://localhost:8080/api/v1/mangas/price/paginated?min=100&max=500&page=1&page_size=5"
```

### **5. Navigation Examples** 🧭

```bash
# First page
curl "http://localhost:8080/api/v1/mangas/paginated?page=1&page_size=10"

# Next page (from pagination.next_page)
curl "http://localhost:8080/api/v1/mangas/paginated?page=2&page_size=10"

# Last page calculation: total_pages from response
curl "http://localhost:8080/api/v1/mangas/paginated?page=5&page_size=10"
```

---

## 🎨 **Frontend Integration**

### **React Example** ⚛️

```typescript
interface PaginationData {
  current_page: number;
  page_size: number;
  total_items: number;
  total_pages: number;
  has_next_page: boolean;
  has_prev_page: boolean;
  next_page?: number;
  previous_page?: number;
}

interface ApiResponse<T> {
  success: boolean;
  message: string;
  data: {
    data: T[];
    pagination: PaginationData;
  };
}

const fetchMangasPage = async (page: number, pageSize: number = 10) => {
  const response = await fetch(
    `http://localhost:8080/api/v1/mangas/paginated?page=${page}&page_size=${pageSize}`
  );
  return response.json() as Promise<ApiResponse<Manga>>;
};

// Usage
const [currentPage, setCurrentPage] = useState(1);
const [mangas, setMangas] = useState([]);
const [pagination, setPagination] = useState<PaginationData | null>(null);

useEffect(() => {
  fetchMangasPage(currentPage).then((response) => {
    setMangas(response.data.data);
    setPagination(response.data.pagination);
  });
}, [currentPage]);
```

### **Pagination Component** 🧩

```typescript
const PaginationComponent = ({ pagination, onPageChange }) => (
  <div className="pagination">
    <button 
      disabled={!pagination.has_prev_page}
      onClick={() => onPageChange(pagination.previous_page)}
    >
      Previous
    </button>
    
    <span>
      Page {pagination.current_page} of {pagination.total_pages}
      ({pagination.total_items} total items)
    </span>
    
    <button 
      disabled={!pagination.has_next_page}
      onClick={() => onPageChange(pagination.next_page)}
    >
      Next
    </button>
  </div>
);
```

---

## ⚡ **Performance Considerations**

### **Database Level**
- ใช้ `OFFSET` และ `LIMIT` สำหรับ pagination
- Database indexes ที่เหมาะสม:
  ```sql
  CREATE INDEX idx_mangas_created_at ON mangas(created_at);
  CREATE INDEX idx_mangas_is_active ON mangas(is_active);
  CREATE INDEX idx_mangas_price ON mangas(price);
  ```

### **Application Level**
- **Count Query Optimization**: แยก count query ออกจาก data query
- **Caching**: Cache total count สำหรับ queries ที่ไม่เปลี่ยนแปลงบ่อย
- **Validation**: จำกัด `page_size` สูงสุด 100 items

### **Best Practices** 🌟
```go
// ✅ Good: Efficient pagination
pagination := domain.NewPaginationRequest(1, 20)
result, err := service.GetMangasPaginated(pagination)

// ❌ Bad: Loading all data
allMangas, err := service.GetMangas() // Don't do this for large datasets
```

---

## 🛡️ **Error Handling**

### **Common Error Scenarios**

| **Error** | **Status** | **Cause** | **Solution** |
|-----------|------------|-----------|--------------|
| Invalid page number | `400` | `page < 1` | Use `page >= 1` |
| Invalid page size | `400` | `page_size < 1 or > 100` | Use `1 <= page_size <= 100` |
| Invalid price range | `400` | `min > max` | Ensure `min <= max` |
| Database error | `500` | DB connection issues | Check server logs |

### **Error Response Format**

```json
{
  "success": false,
  "message": "Invalid pagination parameters",
  "error": {
    "field": "page_size",
    "value": 150,
    "constraint": "max=100"
  }
}
```

---

## 🔄 **Migration from Non-Paginated APIs**

### **Backward Compatibility**

Original endpoints ยังคงใช้งานได้:
```bash
# Non-paginated (existing)
GET /api/v1/mangas              # All mangas
GET /api/v1/mangas/active       # Active mangas
GET /api/v1/mangas/user/4       # User mangas

# Paginated (new)
GET /api/v1/mangas/paginated              # All mangas (paginated)
GET /api/v1/mangas/active/paginated       # Active mangas (paginated)
GET /api/v1/mangas/user/4/paginated       # User mangas (paginated)
```

### **Migration Strategy**

1. **Phase 1**: ใช้ non-paginated APIs ต่อไป
2. **Phase 2**: Gradual migration ไปใช้ paginated APIs
3. **Phase 3**: Deprecate non-paginated APIs (ในอนาคต)

---

## 🧪 **Testing Guide**

### **Quick Test Commands**

```bash
# Basic pagination
curl "http://localhost:8080/api/v1/mangas/paginated?page=1&page_size=5"

# Test navigation
curl "http://localhost:8080/api/v1/mangas/paginated?page=2&page_size=5"

# Test edge cases
curl "http://localhost:8080/api/v1/mangas/paginated?page=999&page_size=10"  # Empty result
curl "http://localhost:8080/api/v1/mangas/paginated?page=0&page_size=150"   # Auto-correction

# Test all endpoints
curl "http://localhost:8080/api/v1/mangas/active/paginated?page=1&page_size=3"
curl "http://localhost:8080/api/v1/mangas/user/4/paginated?page=1&page_size=10"
curl "http://localhost:8080/api/v1/mangas/price/paginated?min=100&max=300&page=1&page_size=5"
```

### **Expected Response Structure**

```json
{
  "success": true,
  "message": "Paginated mangas retrieved successfully",
  "data": {
    "data": [...],
    "pagination": {
      "current_page": 1,
      "page_size": 10,
      "total_items": 25,
      "total_pages": 3,
      "has_next_page": true,
      "has_prev_page": false,
      "next_page": 2
    }
  }
}
```

---

## 🚀 **Future Enhancements**

### **Planned Features**
- **Cursor-based Pagination**: สำหรับ real-time data
- **Sort Parameters**: `sort_by`, `sort_order`
- **Search Integration**: Search + Pagination
- **Cache Layer**: Redis caching สำหรับ popular queries
- **Rate Limiting**: ป้องกัน abuse

### **Advanced Usage**
```bash
# Future: Sorting + Pagination
curl "http://localhost:8080/api/v1/mangas/paginated?page=1&page_size=10&sort_by=price&sort_order=desc"

# Future: Search + Pagination  
curl "http://localhost:8080/api/v1/mangas/search/paginated?q=naruto&page=1&page_size=10"
```

---

## 📝 **Quick Summary**

### **✅ What You Get**
- **Performance**: Efficient data loading with LIMIT/OFFSET
- **User Experience**: Easy navigation with page controls
- **Scalability**: Handles large datasets gracefully
- **Consistency**: Standardized pagination across all endpoints

### **🎯 Key Features**
- **Auto-correction**: Invalid parameters get corrected automatically
- **Type Safety**: Go Generics for type-safe pagination
- **Flexible**: Works with any entity (Manga, User, etc.)
- **Frontend Ready**: Complete metadata for UI pagination components

### **🚀 Ready to Use**
All pagination endpoints are production-ready with:
- ✅ Input validation and error handling
- ✅ Performance optimization (database indexes recommended)
- ✅ Complete documentation and examples
- ✅ Frontend integration examples (React/TypeScript)

---

🎉 **Pagination System Complete!** Now you can handle large datasets efficiently while maintaining excellent user experience. For implementation details, see [How to Add New API](HOW_TO_ADD_NEW_API.md) guide. 🚀 