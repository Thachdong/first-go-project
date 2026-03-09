package repository

import (
	"errors"

	"first-go-project/internal/infrastructure/database/models"
	"first-go-project/internal/modules/user/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id uint) (*domain.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return r.toDomain(&user), nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return r.toDomain(&user), nil
}

func (r *UserRepository) FindByUsername(username string) (*domain.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return r.toDomain(&user), nil
}

func (r *UserRepository) Create(user *domain.User) error {
	dbUser := r.toModel(user)
	if err := r.db.Create(dbUser).Error; err != nil {
		return err
	}
	user.ID = dbUser.ID
	user.CreatedAt = dbUser.CreatedAt
	user.UpdatedAt = dbUser.UpdatedAt
	return nil
}

func (r *UserRepository) Update(user *domain.User) error {
	dbUser := r.toModel(user)
	if err := r.db.Save(dbUser).Error; err != nil {
		return err
	}
	user.UpdatedAt = dbUser.UpdatedAt
	return nil
}

func (r *UserRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) toDomain(user *models.User) *domain.User {
	return &domain.User{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		FullName:  user.FullName,
		Phone:     user.Phone,
		Address:   user.Address,
		City:      user.City,
		Province:  user.Province,
		ZipCode:   user.ZipCode,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (r *UserRepository) toModel(user *domain.User) *models.User {
	return &models.User{
		BaseModel: models.BaseModel{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		FullName: user.FullName,
		Phone:    user.Phone,
		Address:  user.Address,
		City:     user.City,
		Province: user.Province,
		ZipCode:  user.ZipCode,
	}
}
