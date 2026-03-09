# Infrastructure Layer

This directory contains implementations of external dependencies and infrastructure concerns.

## Database

### GORM Setup (`database/postgres.go`)
- PostgreSQL connection with GORM ORM
- Connection pooling configured (10 idle, 100 max open connections)
- Automatic UTC timestamps
- Health check on startup

### Base Model (`database/models/base.go`)
Common fields for all database entities:
- `ID` (primary key)
- `CreatedAt`, `UpdatedAt` (timestamps)
- `DeletedAt` (soft delete support)

## Repository Implementations

Repository implementations go in the `repository/` subdirectory. Pattern:

```go
package repository

import "gorm.io/gorm"

type UserRepositoryImpl struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
    return &UserRepositoryImpl{db: db}
}

// Implement methods defined in domain layer interfaces
func (r *UserRepositoryImpl) FindByID(id uint) (*domain.User, error) {
    var user models.User
    if err := r.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return user.ToDomain(), nil
}
```

## Key Principles
1. Inject `*gorm.DB` into repository constructors
2. Map between GORM models and domain entities
3. Handle GORM errors and convert to domain errors
4. Keep infrastructure details isolated from domain logic
