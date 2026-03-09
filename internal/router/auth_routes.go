package router

import (
	"first-go-project/internal/modules/auth/delivery"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, authHandler *delivery.AuthHandler) {
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/register", authHandler.Register)
	}
}
