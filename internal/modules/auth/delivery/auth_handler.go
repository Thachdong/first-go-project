package delivery

import (
	"errors"
	"net/http"

	"first-go-project/internal/modules/user/domain"
	userusecase "first-go-project/internal/modules/user/usecase"
	"first-go-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userUsecase *userusecase.UserUsecase
}

func NewAuthHandler(userUsecase *userusecase.UserUsecase) *AuthHandler {
	return &AuthHandler{
		userUsecase: userUsecase,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Province string `json:"province"`
	ZipCode  string `json:"zip_code"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Province string `json:"province"`
	ZipCode  string `json:"zip_code"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	newUser := &domain.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Phone:    req.Phone,
		Address:  req.Address,
		City:     req.City,
		Province: req.Province,
		ZipCode:  req.ZipCode,
	}

	if err := h.userUsecase.Register(newUser); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			response.Error(c, http.StatusConflict, "User already exists")
			return
		}
		if errors.Is(err, domain.ErrInvalidUserData) {
			response.Error(c, http.StatusBadRequest, "Invalid user data")
			return
		}
		response.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	registerResponse := RegisterResponse{
		ID:       newUser.ID,
		Email:    newUser.Email,
		Username: newUser.Username,
		FullName: newUser.FullName,
		Phone:    newUser.Phone,
		Address:  newUser.Address,
		City:     newUser.City,
		Province: newUser.Province,
		ZipCode:  newUser.ZipCode,
	}

	response.Success(c, http.StatusCreated, registerResponse, "User registered successfully")
}
