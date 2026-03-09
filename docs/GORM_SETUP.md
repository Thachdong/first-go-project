# GORM + PostgreSQL Quick Reference

## What Was Configured

### 1. Dependencies Added
- `gorm.io/gorm` - GORM ORM
- `gorm.io/driver/postgres` - PostgreSQL driver

### 2. Files Created
- `internal/infrastructure/database/postgres.go` - Database connection setup
- `internal/infrastructure/database/models/base.go` - Base model with common fields
- `internal/infrastructure/README.md` - Infrastructure documentation

### 3. Files Updated
- `cmd/server/main.go` - Added DB initialization and connection management
- `.github/instructions/copilot-instructions.md` - Updated with GORM patterns

## Quick Start

1. **Start PostgreSQL**: 
   ```bash
   docker-compose up -d
   ```

2. **Run the app**:
   ```bash
   air
   # or manually:
   go run cmd/server/main.go
   ```

3. **Add a model** (example for User module):
   ```go
   // internal/infrastructure/database/models/user.go
   package models
   
   type User struct {
       BaseModel
       Email    string `gorm:"uniqueIndex;not null"`
       Username string `gorm:"uniqueIndex;not null"`
       Password string `gorm:"not null"`
   }
   
   func (u *User) TableName() string {
       return "users"
   }
   ```

4. **Register versioned migrations in main.go**:
   ```go
   migrationsPath, err := filepath.Abs("internal/infrastructure/database/migrations")
   if err != nil {
       log.Fatalf("Failed to resolve migrations path: %v", err)
   }

   if err := db.RunMigrations(migrationsPath); err != nil {
       log.Fatalf("Failed to run migrations: %v", err)
   }
   ```

5. **Create repository implementation**:
   ```go
   // internal/infrastructure/repository/user_repository.go
   package repository
   
   import "gorm.io/gorm"
   
   type UserRepository struct {
       db *gorm.DB
   }
   
   func NewUserRepository(db *gorm.DB) *UserRepository {
       return &UserRepository{db: db}
   }
   
   // Implement interface methods...
   ```

## GORM Configuration Details

- **Connection Pool**: 10 idle, 100 max open, 1h lifetime
- **Logging**: Info level (shows SQL queries)
- **Timestamps**: Automatic UTC timestamps
- **Soft Delete**: Enabled via `gorm.DeletedAt` in `BaseModel`

## Versioned Migrations

- Migration files live in `internal/infrastructure/database/migrations`
- Use paired files per version:
    - `000001_description.up.sql`
    - `000001_description.down.sql`
- Applied versions are tracked in PostgreSQL by `golang-migrate` via the `schema_migrations` table
- On app startup, pending migrations are applied automatically

## Environment Variables

Configure via `.env` (see `.env.example`):
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=first_go_db
DB_SSLMODE=disable
```
