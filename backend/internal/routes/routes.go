package routes

import (
	"backend/internal/common"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, dependencies *common.Dependencies) {

	apiGroup := router.Group("/api")

	apiGroup.GET("/auth", dependencies.AuthHandler.Auth)
}
