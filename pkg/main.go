package main

import (
	"binvault/pkg/api"
	"binvault/pkg/cfg"
	"binvault/pkg/clients/filesystem"
	"binvault/pkg/compression"
	"binvault/pkg/database"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

var workers = runtime.NumCPU()

func main() {
	auth := cfg.GetAuth()
	log.Println("Auth enabled", auth.Enabled)
	debug()

	go filesystem.WatchFolder("temp")
	go compression.Init(workers)

	database.Init()
	api.StartServer()

}

func debug() {
	if !cfg.EnvExists("SSH_KEY") {
		return
	}

	key, err := getPrivateKey()
	if err != nil {
		fmt.Printf("failed to get private key: %s\n", err)
		return
	}
	// Build a JWT!
	tok, err := jwt.NewBuilder().
		Issuer(`github.com/lestrrat-go/jwx`).
		IssuedAt(time.Now()).
		Build()
	if err != nil {
		fmt.Printf("failed to build token: %s\n", err)
		return
	}
	signed, err := jwt.Sign(tok, jwa.RS256, key)
	if err != nil {
		log.Fatalf("Failed to sign token: %v", err)
	}
	fmt.Printf("Signed token: %s\n", signed)
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	key := cfg.GetVar("SSH_KEY")

	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, fmt.Errorf("no PEM block found")
	}

	// Try parsing as PKCS1
	if pkcs1Key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return pkcs1Key, nil
	}

	// Otherwise, try parsing as PKCS8
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, err
	}

	return rsaKey, nil
}
