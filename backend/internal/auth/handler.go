package auth

import (
	"backend/internal/common/utils"
	"backend/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service Service
}

type AuthResponse struct {
	Status string    `json:"status"`
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`
}

func (h *Handler) Auth(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	context := ctx.Request.Context()
	token = utils.ParseToken(token)

	user, status, err := h.Service.Authorize(context, token)
	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(status, AuthResponse{
		Status: http.StatusText(status),
		UserID: user.PrivyID,
		User:   *user,
	})
}

func (h *Handler) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "logged out"})
}
