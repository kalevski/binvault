package auth

import (
	"crypto/rsa"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

func GenerateJWT(rsa *rsa.PrivateKey, claims map[string]any) (string, error) {
	key := jose.SigningKey{
		Algorithm: jose.SignatureAlgorithm(jose.RS256),
		Key:       rsa,
	}
	opts := &jose.SignerOptions{}
	opts.WithType("JWT")
	signer, err := jose.NewSigner(key, opts)
	if err != nil {
		return "", err
	}
	jwt, err := jwt.Signed(signer).Claims(claims).Serialize()
	if err != nil {
		return "", err
	}
	return jwt, nil
}
