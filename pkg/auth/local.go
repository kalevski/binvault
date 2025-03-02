package auth

import (
	"binvault/pkg/cfg"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

func LoadPEM() (*rsa.PrivateKey, error) {
	path := filepath.Join(cfg.GetVars().DATA_PATH, cfg.GetVars().PEM_FILENAME)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(data)
	if pemBlock == nil {
		return nil, fmt.Errorf("not valid private PEM key file")
	}
	parsedKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return parsedKey, nil
}

func LoadRSAPublicKey() *rsa.PublicKey {
	data := cfg.GetVars().SSH_PUBLIC_KEY
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
	data := cfg.GetVars().SSH_PRIVATE_KEY
	parsedKey, err := ssh.ParseRawPrivateKey([]byte(data))
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
	rsaPrivKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		log.Fatalf("Key is not an RSA private key")
	}
	return rsaPrivKey
}
