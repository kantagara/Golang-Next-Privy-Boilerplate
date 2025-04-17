package auth

import (
	"backend/internal/user"
	"encoding/json"
	"errors"
	"fmt"
)

type LinkedAccount struct {
	Type             string `json:"type"`
	Address          string `json:"address"`
	ChainType        string `json:"chain_type,omitempty"`
	WalletClientType string `json:"wallet_client_type,omitempty"`
	Lv               int64  `json:"lv"`
}

func CreateUserFromClaims(claims *PrivyClaims) (*user.User, error) {
	var linkedAccounts []LinkedAccount
	err := json.Unmarshal([]byte(claims.LinkedAccounts), &linkedAccounts)
	if err != nil {
		return nil, fmt.Errorf("failed to parse linked_accounts: %w", err)
	}

	var walletAccount *LinkedAccount
	for _, acc := range linkedAccounts {
		if acc.Type == "wallet" {
			walletAccount = &acc
			break
		}
	}

	if walletAccount == nil {
		return nil, errors.New("no wallet account found")
	}

	newUser := &user.User{
		PrivyID:       claims.UserId,
		WalletAddress: walletAccount.Address,
		Username:      "",
	}

	return newUser, nil
}
