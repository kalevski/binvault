package auth

import (
	"crypto/rsa"
	"fmt"

	"gopkg.in/square/go-jose.v2/jwt"
)

func ValidateJWT(key *rsa.PublicKey, payload string) (map[string]any, error) {
	token, err := jwt.ParseSigned(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %w", err)
	}

	claims := make(map[string]interface{})

	if err := token.Claims(key, &claims); err != nil {
		return nil, fmt.Errorf("failed to validate JWT: %w", err)
	}

	return claims, nil
}
