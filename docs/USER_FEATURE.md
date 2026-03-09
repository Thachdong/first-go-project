# User Feature Test Guide

## API Endpoint

**Get User Profile**: `GET /api/v1/users/:id`

## Testing Steps

### 1. Start Database
```bash
docker-compose up -d
```

### 2. Run the Server
```bash
air
# or
go run cmd/server/main.go
```

### 3. Seed Test Data
Open a new terminal and run:
```bash
go run scripts/seed.go
```

This will create 3 test users:
- ID 1: john.doe@example.com (John Doe) - New York, NY
- ID 2: jane.smith@example.com (Jane Smith) - Los Angeles, CA
- ID 3: bob.wilson@example.com (Bob Wilson) - Chicago, IL

### 4. Test the API

**Using curl:**
```bash
# Get user with ID 1
curl http://localhost:8080/api/v1/users/1

# Get user with ID 2
curl http://localhost:8080/api/v1/users/2

# Test with invalid ID
curl http://localhost:8080/api/v1/users/999
```

**Expected Response (Success):**
```json
{
  "status": 200,
  "data": {
    "id": 1,
    "email": "john.doe@example.com",
    "username": "johndoe",
    "full_name": "John Doe",
    "phone": "+1234567890",
    "address": "123 Main Street",
    "city": "New York",
    "province": "NY",
    "zip_code": "10001"
  },
  "message": "User profile retrieved successfully"
}
```

**Expected Response (Not Found):**
```json
{
  "status": 404,
  "message": "User not found"
}
```

**Expected Response (Bad Request):**
```json
{
  "status": 400,
  "message": "Invalid user ID"
}
```

## Architecture Overview

The user feature follows Clean Architecture:

```
┌─────────────────────────────────────────────────────┐
│ Delivery Layer (HTTP Handler)                       │
│ internal/modules/user/delivery/user_handler.go      │
│ - Handles HTTP requests/responses                   │
│ - Converts DTOs                                      │
│ - Maps domain errors to HTTP status codes           │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│ Usecase Layer (Business Logic)                      │
│ internal/modules/user/usecase/user_usecase.go       │
│ - Orchestrates business logic                       │
│ - Uses domain interfaces                            │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│ Domain Layer (Entities & Interfaces)                │
│ internal/modules/user/domain/                       │
│ - User entity                                        │
│ - Repository interface                              │
│ - Domain errors                                      │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│ Infrastructure Layer (Repository Implementation)    │
│ internal/infrastructure/repository/user_repository.go│
│ - Implements UserRepository interface               │
│ - GORM database operations                          │
│ - Model ↔ Entity conversion                         │
└─────────────────────────────────────────────────────┘
```

## Files Created

### Domain Layer
- `internal/modules/user/domain/user.go` - User entity & repository interface
- `internal/modules/user/domain/errors.go` - Domain-specific errors

### Database Models
- `internal/infrastructure/database/models/user.go` - GORM model

### Repository
- `internal/infrastructure/repository/user_repository.go` - Repository implementation

### Usecase
- `internal/modules/user/usecase/user_usecase.go` - Business logic

### Delivery
- `internal/modules/user/delivery/user_handler.go` - HTTP handlers

### Router
- `internal/router/user_routes.go` - Route definitions

### Scripts
- `scripts/seed.go` - Database seeding script

## Next Steps

You can extend this feature by adding:
1. Create user endpoint (POST /api/v1/users)
2. Update user endpoint (PUT /api/v1/users/:id)
3. Delete user endpoint (DELETE /api/v1/users/:id)
4. List users endpoint (GET /api/v1/users)
5. Authentication & password hashing
6. Input validation
7. Unit tests for each layer
