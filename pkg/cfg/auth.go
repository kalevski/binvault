package cfg

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/lestrrat-go/jwx/jwk"
	"golang.org/x/crypto/ssh"
)

type Auth struct {
	Enabled   bool
	Jwk       *jwk.Set
	PayloadId string
}

type RSAKey struct {
	N *big.Int
	E int
}

var auth *Auth
var initialized bool = false

func GetAuth() *Auth {

	if !EnvExists("JWT_PAYLOAD_ID") {
		panic("environment variable JWT_PAYLOAD_ID is not set")
	}

	if initialized {
		return getInstance()
	}

	if EnvExists("JWKS_URL") && EnvExists("JWKS_KID") {
		log.Println("JWT auth enabled (JWKS)")
		auth = &Auth{
			Enabled:   true,
			Jwk:       nil,
			PayloadId: GetVar("JWT_PAYLOAD_ID"),
		}
		go startJWKSUpdater()
	} else if EnvExists("SSH_PUBKEY") {
		log.Println("JWT auth enabled (SSH KEY)")
		pubkey := GetVar("SSH_PUBKEY")
		set := jwk.NewSet()

		pem, err := convertSSHRsaPublicKeyToPEM([]byte(pubkey))
		if err != nil {
			log.Fatalf("failed to convert SSH RSA public key to PEM: %s", err)
		}

		key, err := jwk.ParseKey(pem, jwk.WithPEM(true))
		if err != nil {
			log.Fatalf("failed to parse PEM encoded key: %s", err)
		}

		if err := key.Set(jwk.KeyIDKey, GetVar("JWKS_KID")); err != nil {
			log.Fatalf("failed to set key ID: %s", err)
		}
		set.Add(key)

		auth = &Auth{
			Enabled:   true,
			Jwk:       &set,
			PayloadId: GetVar("JWT_PAYLOAD_ID"),
		}
	} else {
		log.Println("JWT auth disabled")
		auth = &Auth{
			Enabled:   false,
			Jwk:       nil,
			PayloadId: "",
		}
	}

	return getInstance()
}

func getInstance() *Auth {
	return auth
}

func startJWKSUpdater() {
	for {
		updateJWKS()
		time.Sleep(15 * time.Minute)
	}
}

func updateJWKS() {
	jwksURL := GetVar("JWKS_URL")
	set, err := jwk.Fetch(context.Background(), jwksURL)
	if err != nil {
		log.Printf("failed to parse JWK: %s", err)
		return
	}
	auth.Jwk = &set
}

func convertSSHRsaPublicKeyToPEM(sshKeyBytes []byte) ([]byte, error) {
	pubKey, _, _, _, err := ssh.ParseAuthorizedKey(sshKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SSH public key: %w", err)
	}
	cryptoPubKey, ok := pubKey.(ssh.CryptoPublicKey)
	if !ok {
		return nil, errors.New("SSH key does not implement CryptoPublicKey")
	}
	rsaPubKey, ok := cryptoPubKey.CryptoPublicKey().(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("key is not an RSA public key")
	}
	derBytes, err := x509.MarshalPKIXPublicKey(rsaPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal RSA public key: %w", err)
	}
	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	})
	return pemBytes, nil
}
