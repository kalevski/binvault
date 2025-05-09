package auth

import (
	"binvault/pkg/env"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

type PEM struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func LoadPEM() *PEM {
	privatePath := filepath.Join(env.GetVars().DATA_PATH, env.GetVars().PEM_PRIVATE_FILENAME)
	publicPath := filepath.Join(env.GetVars().DATA_PATH, env.GetVars().PEM_PUBLIC_FILENAME)

	privateData, err := os.ReadFile(privatePath)
	if err != nil {
		return nil
	}

	publicData, err := os.ReadFile(publicPath)
	if err != nil {
		return nil
	}

	pemBlock, _ := pem.Decode(privateData)
	if pemBlock == nil {
		return nil
	}
	parsedKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil
	}

	pubPremBlock, _ := pem.Decode(publicData)
	if pubPremBlock == nil {
		return nil
	}

	parsedPubKey, err := x509.ParsePKCS1PublicKey(pubPremBlock.Bytes)
	if err != nil {
		return nil
	}
	return &PEM{
		PrivateKey: parsedKey,
		PublicKey:  parsedPubKey,
	}
}

func LoadRSAPublicKey() *rsa.PublicKey {
	data := env.GetVars().RSA_PUBLIC_KEY
	parsedKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(data))
	if err != nil {
		log.Fatalf("Failed to parse public key: %v", err)
	}
	cryptoPubKey, ok := parsedKey.(ssh.CryptoPublicKey)
	if !ok {
		log.Fatalf("Parsed key does not implement CryptoPublicKey")
	}
	pub := cryptoPubKey.CryptoPublicKey()

	rsa, ok := pub.(*rsa.PublicKey)
	if !ok {
		log.Fatalf("Key is not an RSA public key")
	}

	return rsa
}

func LoadRSAPrivateKey() *rsa.PrivateKey {
	if !env.EnvExists("RSA_PRIVATE_KEY") {
		return nil
	}

	data := env.GetVars().RSA_PRIVATE_KEY
	parsedKey, err := ssh.ParseRawPrivateKey([]byte(data))
	if err != nil {
		return nil
	}
	rsaPrivKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		log.Println("Key is not an RSA private key")
		return nil
	}
	return rsaPrivKey
}
