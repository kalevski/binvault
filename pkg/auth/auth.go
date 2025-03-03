package auth

import (
	"binvault/pkg/cfg"
	"crypto/rsa"
	"log"
	"time"
)

type Auth struct {
	Enabled    bool
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	ClaimId    string
}

var auth *Auth
var initialized bool = false

func GetAuth() *Auth {

	if initialized {
		return getInstance()
	}

	var privateKey *rsa.PrivateKey = nil
	pem := LoadPEM()
	if pem != nil {
		privateKey = pem.PrivateKey
	}

	if privateKey == nil {
		privateKey = LoadRSAPrivateKey()
	}

	auth = &Auth{
		Enabled:    false,
		PrivateKey: privateKey,
	}

	if !cfg.EnvExists("JWT_CLAIM_ID") {
		panic("environment variable JWT_CLAIM_ID is not set")
	} else {
		auth.ClaimId = cfg.GetVar("JWT_CLAIM_ID")
	}

	if cfg.EnvExists("JWKS_URL") && cfg.EnvExists("JWKS_KID") {
		log.Println("JWT auth enabled (JWKS)")
		auth.Enabled = true
		go updateJWKS()
	} else if cfg.EnvExists("SSH_PUBLIC_KEY") {
		auth.Enabled = true
		auth.PublicKey = LoadRSAPublicKey()
	} else if pem != nil {
		auth.Enabled = true
		auth.PublicKey = pem.PublicKey
	}
	return getInstance()
}

func getInstance() *Auth {
	return auth
}

func updateJWKS() {
	jwksUrl := cfg.GetVar("JWKS_URL")
	kid := cfg.GetVar("JWKS_KID")
	for {
		key, err := LoadRSAPublicKeyFromJWKS(jwksUrl, kid)
		if err != nil {
			log.Printf("failed to load RSA public key from JWKS: %v", err)
		} else {
			log.Printf("JWKS [%s] updated", kid)
			auth.PublicKey = key
		}
		time.Sleep(5 * time.Minute)
	}
}
