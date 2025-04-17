package auth

import (
	"backend/internal/common/utils"
	"backend/internal/user"
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	VerifyUser(token string) (*PrivyClaims, error)
	VerifyIdentityToken(token string) (*PrivyClaims, error)

	Authorize(ctx context.Context, token string) (*user.User, int, error)
}

func NewService(repo user.Repository) Service {
	return &authService{
		repo: repo,
	}
}

type authService struct {
	repo user.Repository
}

func (a *authService) Authorize(ctx context.Context, token string) (*user.User, int, error) {

	claims, err := a.VerifyIdentityToken(token)

	if err != nil {
		return nil, http.StatusUnauthorized, err
	}

	foundUser, err := a.repo.GetUserById(ctx, claims.UserId)
	if err == nil && foundUser != nil {
		return foundUser, http.StatusOK, nil
	}

	newUser, err := CreateUserFromClaims(claims)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to create user from claims: %w", err)
	}

	savedUser, err := a.repo.Create(ctx, newUser)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to save user: %w", err)
	}

	return savedUser, http.StatusCreated, nil // User Created
}

func (a *authService) VerifyIdentityToken(tokenString string) (*PrivyClaims, error) {

	key := os.Getenv("PRIVY_VERIFICATION_KEY")
	key = strings.ReplaceAll(key, "\\n", "\n")

	block, _ := pem.Decode([]byte(key))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key format")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not ECDSA")
	}

	claims := &PrivyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodES256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return ecdsaPubKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid JWT: " + err.Error())
	}

	if claims.Issuer != "privy.io" {
		return nil, errors.New("invalid issuer")
	}

	if claims.AppId != os.Getenv("PRIVY_APP_ID") {
		return nil, errors.New("invalid audience")
	}

	return claims, nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if token.Method.Alg() != "ES256" {
		return nil, fmt.Errorf("unexpected JWT signing method: %v", token.Header["alg"])
	}

	key := os.Getenv("PRIVY_VERIFICATION_KEY")
	key = strings.ReplaceAll(key, "\\n", "\n")

	pubKey, err := jwt.ParseECPublicKeyFromPEM([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	return pubKey, nil
}

func (a *authService) VerifyUser(rawToken string) (*PrivyClaims, error) {
	rawToken = utils.ParseToken(rawToken)
	token, err := jwt.ParseWithClaims(rawToken, &PrivyClaims{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT signature: %w", err)
	}

	claims, ok := token.Claims.(*PrivyClaims)
	if !ok {
		return nil, errors.New("JWT does not contain PrivyClaims")
	}

	if err := claims.Valid(); err != nil {
		return nil, fmt.Errorf("invalid JWT claims: %w", err)
	}

	return claims, nil
}
