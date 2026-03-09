package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"first-go-project/internal/modules/user/domain"
	"first-go-project/internal/modules/user/usecase"
	"first-go-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

type UserResponse struct {
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

func (h *UserHandler) GetProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.userUsecase.GetUserProfile(uint(id))
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			response.Error(c, http.StatusNotFound, "User not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	userResponse := UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		Phone:    user.Phone,
		Address:  user.Address,
		City:     user.City,
		Province: user.Province,
		ZipCode:  user.ZipCode,
	}

	response.Success(c, http.StatusOK, userResponse, "User profile retrieved successfully")
}
