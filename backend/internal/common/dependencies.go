package common

import (
	"backend/internal/auth"
	"backend/internal/user"
)

type Dependencies struct {
	UserHandler *user.Handler
	AuthHandler *auth.Handler
}
