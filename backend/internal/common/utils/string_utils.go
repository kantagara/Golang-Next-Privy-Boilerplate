package utils

import (
	"strings"
)

// ParseToken Removes Bearer from Auth
func ParseToken(token string) string {
	parts := strings.Split(token, " ")

	return parts[1]
}
