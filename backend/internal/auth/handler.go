package auth

import (
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

	user, status, err := h.Service.Authorize(ctx)
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
