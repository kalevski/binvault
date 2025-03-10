package auth

import (
	"crypto/rsa"
	"fmt"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

var signatureAlgorithms []jose.SignatureAlgorithm = []jose.SignatureAlgorithm{
	jose.RS256,
}

func ValidateJWT(key *rsa.PublicKey, token string) (map[string]any, error) {

	webToken, err := jwt.ParseSigned(token, signatureAlgorithms)
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
	webToken, err := jwt.ParseSigned(token, signatureAlgorithms)
	if err != nil {
		return nil
	}

	claims := make(map[string]any)

	if err := webToken.UnsafeClaimsWithoutVerification(&claims); err != nil {
		return nil
	}
	return claims
}
