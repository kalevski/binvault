package auth

import (
	"crypto/rsa"
	"fmt"

	"gopkg.in/square/go-jose.v2/jwt"
)

func ValidateJWT(key *rsa.PublicKey, token string) (map[string]any, error) {
	webToken, err := jwt.ParseSigned(token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %w", err)
	}

	claims := make(map[string]interface{})

	if err := webToken.Claims(key, &claims); err != nil {
		return nil, fmt.Errorf("failed to validate JWT: %w", err)
	}

	return claims, nil
}

func GetClaims(token string) map[string]any {
	webToken, err := jwt.ParseSigned(token)
	if err != nil {
		return nil
	}

	claims := make(map[string]any)

	if err := webToken.UnsafeClaimsWithoutVerification(&claims); err != nil {
		return nil
	}
	return claims
}
