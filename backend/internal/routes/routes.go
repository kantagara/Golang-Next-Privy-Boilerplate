package routes

import (
	"backend/internal/auth"
	"backend/internal/common"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, dependencies *common.Dependencies) {

	apiGroup := router.Group("/api")
	apiGroup.Use(auth.AuthMiddleware(dependencies.AuthHandler.Service))

	apiGroup.GET("/auth", dependencies.AuthHandler.Auth)
}
