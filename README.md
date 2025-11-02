# 📘 Golang REST API — Gin + PostgreSQL + JWT

A simple RESTful API built with **Golang**, **Gin**, and **PostgreSQL**, featuring:
- JWT Authentication  
- User, Book, and Category Management  
- SQL Migrations using [`rubenv/sql-migrate`](https://github.com/rubenv/sql-migrate)  
- Clean project structure (Controllers, Models, Routes, Middleware, Helpers, Config, utils, migrations)

---

## 🚀 Project Setup

### 1️⃣ Clone the repository
```bash
git clone https://github.com/MinYuliadi/bookstore.git
cd your-repo-name
```

### 2️⃣ Install dependencies
```bash
go mod tidy
```

### 3️⃣ Create .env file
```.env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=mydb
JWT_SECRET=supersecretkey
```

### 🧱 Project Structure

```
├── config/             # Database connection setup
├── controllers/        # Business logic (Auth, Books, Categories)
│   ├── auth/
│   ├── books/
│   └── categories/
├── helpers/            # Constants, utilities, and JWT helpers
├── middleware/         # Auth middleware, method validation
├── migrations/         # SQL migration files
├── models/             # Struct definitions for DB entities
├── routers/            # Route definitions
├── utils/              # Function utilities
├── main.go             # Entry point
└── go.mod
```

### 🧩 API Endpoints

### 🔐 Auth
| Method | Endpoint    | Description                    |
| ------ | ----------- | ------------------------------ |
| `POST` | `/register` | Create a new user              |
| `POST` | `/login`    | User login (returns JWT token) |

/register 
```
curl --location 'http://localhost:8080/api/users/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "yourname",
    "password": "yourpassword"
}'
```

/login 
```
curl --location 'http://localhost:8080/api/users/login' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "username": "yourname",
    "password": "yourpassword"
}'
```

### 📚 Books
| Method   | Endpoint     | Description     |
| -------- | ------------ | --------------- |
| `GET`    | `/books`     | Get all books   |
| `GET`    | `/books/:id` | Get book by ID  |
| `POST`   | `/books`     | Create new book |
| `PATCH`  | `/books/:id` | Update book     |
| `DELETE` | `/books/:id` | Delete book     |

/books get 
```
curl --location 'http://localhost:8080/api/books' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMDk2NTIsImlhdCI6MTc2MjEwNjA1Mn0.5hiJSf1U_jJYy_Ma_FPNKXA8JHa3tyExjO0LsgpgcLs' \
--data ''
```

/books/:id get 
```
curl --location 'http://localhost:8080/api/books/3' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMTcyMTYsImlhdCI6MTc2MjExMzYxNn0.Eamys3gMPZ9GCyhHnosJtgtaRunruhHarBLU8MAay_g' \
--data ''
```

/books post 
```
curl --location 'http://localhost:8080/api/books' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMDk2NTIsImlhdCI6MTc2MjEwNjA1Mn0.5hiJSf1U_jJYy_Ma_FPNKXA8JHa3tyExjO0LsgpgcLs' \
--data '{
    "title": "thin history",
    "description": "A handbook of agile software craftsmanship",
    "image_url": "https://example.com/images/clean-code.jpg",
    "release_year": 2008,
    "price": 250000,
    "total_page": 464,
    "category_id": 5
}'
```

/books/:id patch 
```
curl --location --request PATCH 'http://localhost:8080/api/books/5' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMTcyMTYsImlhdCI6MTc2MjExMzYxNn0.Eamys3gMPZ9GCyhHnosJtgtaRunruhHarBLU8MAay_g' \
--data '{
    "title": "thin history",
    "description": "A handbook of agile software craftsmanship",
    "image_url": "https://example.com/images/clean-code.jpg",
    "release_year": 2024,
    "price": 250000,
    "total_page": 80,
    "category_id": 5
}'
```

/books/:id delete 
```
curl --location --request DELETE 'http://localhost:8080/api/books/3' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMTcyMTYsImlhdCI6MTc2MjExMzYxNn0.Eamys3gMPZ9GCyhHnosJtgtaRunruhHarBLU8MAay_g'
```

### 🏷 Categories
| Method   | Endpoint                 | Description            |
| -------- | ------------------------ | ---------------------- |
| `GET`    | `/categories`            | Get all categories     |
| `GET`    | `/categories/:id`        | Get category by ID     |
| `POST`   | `/categories`            | Create new categories  |
| `PATCH`  | `/categories/:id`        | Update category        |
| `DELETE` | `/categories/:id`        | Delete category        |
| `GET`    | `/categories/:id/books`  | Get books by category  |

/categories get 
```
curl --location 'http://localhost:8080/api/categories' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMTcyMTYsImlhdCI6MTc2MjExMzYxNn0.Eamys3gMPZ9GCyhHnosJtgtaRunruhHarBLU8MAay_g' \
--data ''
```

/categories/:id get 
```
curl --location 'http://localhost:8080/api/categories/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMTcyMTYsImlhdCI6MTc2MjExMzYxNn0.Eamys3gMPZ9GCyhHnosJtgtaRunruhHarBLU8MAay_g' \
--data ''
```

/categories post 
```
curl --location 'http://localhost:8080/api/categories' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMDk2NTIsImlhdCI6MTc2MjEwNjA1Mn0.5hiJSf1U_jJYy_Ma_FPNKXA8JHa3tyExjO0LsgpgcLs' \
--data '{
    "name": "Teknologi"
}'
```

/categories/:id patch 
```
curl --location --request PATCH 'http://localhost:8080/api/categories/4' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMDk2NTIsImlhdCI6MTc2MjEwNjA1Mn0.5hiJSf1U_jJYy_Ma_FPNKXA8JHa3tyExjO0LsgpgcLs' \
--data '{
    "name": "Ilmiah"
}'
```

/categories/:id delete
```
curl --location --request DELETE 'http://localhost:8080/api/categories/7' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMDk2NTIsImlhdCI6MTc2MjEwNjA1Mn0.5hiJSf1U_jJYy_Ma_FPNKXA8JHa3tyExjO0LsgpgcLs' \
--data ''
```

/categories/:id/books get 
```
curl --location 'http://localhost:8080/api/categories/5/books' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkZhbm4iLCJleHAiOjE3NjIxMTcyMTYsImlhdCI6MTc2MjExMzYxNn0.Eamys3gMPZ9GCyhHnosJtgtaRunruhHarBLU8MAay_g'
```

### 🧰 Run the Server
```bash
go run main.go
```

### 🔒 Authentication
All protected routes require a valid JWT Token in the request header:
```
Authorization: Bearer <your_token>
```

### 🧑‍💻 Author

Irfan Nugraha
Front-End & Golang Developer
Made with ❤️ and ☕ power.
