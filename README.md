# Go Course Management REST API

A RESTful API built with Go (Golang) for managing courses, users, and enrollments. This project implements clean architecture principles and provides CRUD operations through HTTP endpoints.

## 🛠 Tech Stack

- **Go** v1.23.4
- **PostgreSQL** - Database
- **GORM** - ORM for database operations
- **Gorilla Mux** - HTTP router and URL matcher
- **godotenv** - Environment configuration

## 📁 Project Structure

```
.
├── internal/
│   ├── course/        # Course domain logic
│   ├── domain/        # Domain entities
│   ├── enrollment/    # Enrollment domain logic
│   └── user/          # User domain logic
├── pkg/
│   ├── bootstrap/     # App initialization
│   └── meta/          # Pagination utilities
└── main.go           # Application entry point
```

## 🗄️ Database Schema

### Users Table

| Field      | Type     |
| ---------- | -------- |
| ID         | UUID     |
| FirstName  | String   |
| LastName   | String   |
| Email      | String   |
| Phone      | String   |
| Timestamps | DateTime |

### Courses Table

| Field      | Type     |
| ---------- | -------- |
| ID         | UUID     |
| Name       | String   |
| StartDate  | DateTime |
| EndDate    | DateTime |
| Timestamps | DateTime |

### Enrollments Table

| Field      | Type      |
| ---------- | --------- |
| ID         | UUID      |
| UserID     | UUID (FK) |
| CourseID   | UUID (FK) |
| Status     | String    |
| Timestamps | DateTime  |

## 🚀 API Endpoints

### Users

```http
GET    /users          # List all users
POST   /users          # Create new user
GET    /users/{id}     # Get user by ID
PATCH  /users/{id}     # Update user
DELETE /users/{id}     # Delete user
```

### Courses

```http
GET    /courses        # List all courses
POST   /courses        # Create new course
GET    /courses/{id}   # Get course by ID
PATCH  /courses/{id}   # Update course
DELETE /courses/{id}   # Delete course
```

### Enrollments

```http
GET    /enrollments    # List all enrollments
POST   /enrollments    # Create new enrollment
```

## 🚀 Getting Started

```bash
go run main.go
```

The server will start at `http://localhost:8000`

## 🔍 Features

- Clean Architecture
- CRUD operations for users and courses
- Course enrollment management
- Pagination support
- Environment-based configuration
- UUID for entity IDs
- Soft delete support
- Timestamp tracking (created_at, updated_at)

## 📘 API Documentation

All endpoints return JSON responses with the following structure:

```json
{
  "status": 200,
  "data": {},
  "error": "",
  "meta": {
    "total_count": 0,
    "page": 1,
    "perPage": 10
  }
}
```

## 📝 Request Examples

### Create User

```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "phone": "1234567890"
}
```

### Create Course

```json
{
  "name": "Introduction to Go",
  "start_date": "2024-03-01",
  "end_date": "2024-06-30"
}
```

### Create Enrollment

```json
{
  "user_id": "uuid-here",
  "course_id": "uuid-here"
}
```
