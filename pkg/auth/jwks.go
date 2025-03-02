package auth

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gopkg.in/square/go-jose.v2"
)

func GetRSAPublicKeyFromJWKS(jwksURL, kid string) (*rsa.PublicKey, error) {
	// Fetch the JWKS JSON from the provided URL.
	resp, err := http.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read JWKS response: %w", err)
	}

	// Unmarshal the JSON into a jose.JSONWebKeySet.
	var jwks jose.JSONWebKeySet
	if err := json.Unmarshal(body, &jwks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JWKS: %w", err)
	}

	// Retrieve the keys that match the provided key ID.
	keys := jwks.Key(kid)
	if len(keys) == 0 {
		return nil, fmt.Errorf("no key found for kid: %s", kid)
	}

	// Pick the first matching key.
	jwk := keys[0]

	// Type-assert that the key is an *rsa.PublicKey.
	rsaPubKey, ok := jwk.Key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key with kid %s is not an RSA public key", kid)
	}

	return rsaPubKey, nil
}
