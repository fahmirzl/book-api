# 📚 Book API

Book API is a simple RESTful API built with Golang and Gin, used to manage data for books, and categories.

---

## 🚀 Feature

- CRUD operations for Books, Categories, and Users  
- Input validation  
- Connection to PostgreSQL  
- Structured JSON responses  
- Middleware for basic authentication  

---

## 🛠️ Tech Stack

- [Go (Golang)](https://golang.org/)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- PostgreSQL
- godotenv

---

## ⚙️ Instalation

1. Clone repository:
   ```bash
   git clone https://github.com/username/book-api.git
   cd book-api
   ```

3. Install dependency:
   ```bash
   go mod tidy
   ```

4. Run API:
   ```bash
   go run main.go
   ```

---

## 📘 Api EndPoint
### 📚 Book

| Method | Endpoint                    | Deskripsi               |
|--------|-----------------------------|-------------------------|
| GET    | `/api/books`                | Retrieve all book       |
| GET    | `/api/books/:id`            | Retrieve book by ID     |
| POST   | `/api/books`                | Add a new book          |
| PUT    | `/api/books/:id`            | Update a book by ID     |
| DELETE | `/api/books/:id`            | Delete a book by ID     |

Example JSON Body:
```json
{
    "title": "Learn Math",
    "description": "Easy learn math.",
    "image_url": "https://example.com/image.jpg",
    "release_year": 2023,
    "price": 100000,
    "total_page": 3000,
    "category_id": 4
}
```

---

### 🗂️ Category

| Method | Endpoint                    | Deskripsi               |
|--------|-----------------------------|-------------------------|
| GET    | `/api/categories`           | Retrieve all category   |
| GET    | `/api/categories/:id`       | Retrieve category by ID |
| POST   | `/api/categories`           | Add a new category      |
| PUT    | `/api/categories/:id`       | Update a category by ID |
| DELETE | `/api/categories/:id`       | Delete a category by ID |

Example JSON Body:
```json
{
    "name": "Myth"
}
```

---

## 🔐 Basic Auth

Use **Basic Authentication** in header:

```http
Authorization: Basic base64(username:password)
```

---

## 🗂 Directory Structure

```
├── handlers/
├── middlewares/
├── db/
├── routers/
├── repositories/
└── structs/
```
