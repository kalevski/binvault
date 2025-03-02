package auth

import (
	"binvault/pkg/cfg"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
)

func GeneratePEM(path string) error {
	pemFilename := cfg.GetVar("PEM_FILENAME")
	file := filepath.Join(path, pemFilename)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	err = pem.Encode(f, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	return err
}
