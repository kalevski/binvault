package auth

import (
	"binvault/pkg/env"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
)

func GeneratePEM(path string) error {
	pemPrivate := env.GetVars().PEM_PRIVATE_FILENAME
	pemPublic := env.GetVars().PEM_PUBLIC_FILENAME

	filePrivate := filepath.Join(path, pemPrivate)
	fp, err := os.Create(filePrivate)
	if err != nil {
		return err
	}
	defer fp.Close()

	filePublic := filepath.Join(path, pemPublic)
	fpub, err := os.Create(filePublic)
	if err != nil {
		return err
	}
	defer fpub.Close()

	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	if err = pem.Encode(fp, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}); err != nil {
		return err
	}

	if err = pem.Encode(fpub, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
	}); err != nil {
		return err
	}

	return err
}
