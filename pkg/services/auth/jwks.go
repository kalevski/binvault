package auth

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-jose/go-jose/v4"
)

func LoadRSAPublicKeyFromJWKS(jwksURL, kid string) (*rsa.PublicKey, error) {
	resp, err := http.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read JWKS response: %w", err)
	}

	var jwks jose.JSONWebKeySet
	if err := json.Unmarshal(body, &jwks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JWKS: %w", err)
	}

	keys := jwks.Key(kid)
	if len(keys) == 0 {
		return nil, fmt.Errorf("no key found for kid: %s", kid)
	}

	jwk := keys[0]

	rsaPubKey, ok := jwk.Key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key with kid %s is not an RSA public key", kid)
	}

	return rsaPubKey, nil
}
