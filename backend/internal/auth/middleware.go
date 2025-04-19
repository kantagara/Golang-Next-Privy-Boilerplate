package auth

import (
	"backend/internal/common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(service AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		token = utils.ParseToken(token)
		claims, err := service.VerifyIdentityToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
