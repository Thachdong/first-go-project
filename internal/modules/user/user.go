package user

import (
	"first-go-project/internal/infrastructure/repository"
	"first-go-project/internal/modules/user/delivery"
	"first-go-project/internal/modules/user/usecase"

	"gorm.io/gorm"
)

func NewHandler(db *gorm.DB) *delivery.UserHandler {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	return delivery.NewUserHandler(userUsecase)
}
