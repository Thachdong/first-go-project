package auth

import (
	"first-go-project/internal/modules/auth/delivery"
	"first-go-project/internal/modules/user/repository"
	"first-go-project/internal/modules/user/usecase"

	"gorm.io/gorm"
)

func NewHandler(db *gorm.DB) *delivery.AuthHandler {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	return delivery.NewAuthHandler(userUsecase)
}
