package router

import (
	"first-go-project/internal/modules/user/delivery"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userHandler *delivery.UserHandler) {
	userGroup := router.Group("/api/v1/users")
	{
		userGroup.GET("/:id", userHandler.GetProfile)
	}
}
