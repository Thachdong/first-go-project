package usecase

import (
	"errors"
	"strings"

	"first-go-project/internal/modules/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) GetUserProfile(id uint) (*domain.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) GetUserByEmail(email string) (*domain.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) Register(user *domain.User) error {
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return domain.ErrInvalidUserData
	}

	if _, err := u.userRepo.FindByEmail(user.Email); err == nil {
		return domain.ErrUserAlreadyExists
	} else if !errors.Is(err, domain.ErrUserNotFound) {
		return err
	}

	if _, err := u.userRepo.FindByUsername(user.Username); err == nil {
		return domain.ErrUserAlreadyExists
	} else if !errors.Is(err, domain.ErrUserNotFound) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err := u.userRepo.Create(user); err != nil {
		return err
	}

	return nil
}
