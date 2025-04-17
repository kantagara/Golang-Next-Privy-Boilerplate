package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type PrivyClaims struct {
	AppId          string `json:"aud,omitempty"`
	LinkedAccounts string `json:"linked_accounts,omitempty"`
	Expiration     uint64 `json:"exp,omitempty"`
	IssuedAt       uint64 `json:"iat,omitempty"`
	Issuer         string `json:"iss,omitempty"`
	UserId         string `json:"sub,omitempty"`
}

func (c *PrivyClaims) Valid() error {
	appId := os.Getenv("PRIVY_APP_ID")

	if c.AppId != appId {
		return errors.New("aud claim must be your Privy App ID")
	}
	if c.Issuer != "privy.io" {
		return errors.New("iss claim must be 'privy.io'")
	}
	if c.Expiration < uint64(time.Now().Unix()) {
		return errors.New("token is expired")
	}
	return nil
}

func (c *PrivyClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(int64(c.Expiration), 0)), nil
}

func (c *PrivyClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	if c.IssuedAt == 0 {
		return nil, nil
	}
	return jwt.NewNumericDate(time.Unix(int64(c.IssuedAt), 0)), nil
}

func (c *PrivyClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (c *PrivyClaims) GetIssuer() (string, error) {
	return c.Issuer, nil
}

func (c *PrivyClaims) GetSubject() (string, error) {
	return c.UserId, nil
}

func (c *PrivyClaims) GetAudience() (jwt.ClaimStrings, error) {
	if c.AppId == "" {
		return nil, nil
	}
	return jwt.ClaimStrings{c.AppId}, nil
}
