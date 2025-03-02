package jwtgen

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var JWTGen = &cobra.Command{
	Use:   "jwt",
	Short: "Generate JWT",
	RunE:  RunJWTGen,
}

var keyFilePath string
var keyAlgorythm string
var jwtAudience string
var claimUserID int64
var claimUserEmail string

func init() {
	flags := JWTGen.Flags()
	flags.StringVar(&keyFilePath, "key-file", ".dev/key.pem", "Path to private key file")
	flags.StringVar(&keyAlgorythm, "alg", "RS256", "Hash algorithm")
	flags.StringVar(&jwtAudience, "aud", "webgame-cloud-api-clients", "JWT Audience claim")
	flags.Int64Var(&claimUserID, "user-id", 0, "Generate JWT for this user")
	flags.StringVar(&claimUserEmail, "email", "", "JWT user email")
}

func RunJWTGen(cmd *cobra.Command, args []string) error {
	if claimUserID == 0 {
		return fmt.Errorf("plase set user-id")
	}
	if claimUserEmail == "" {
		return fmt.Errorf("plase set user email")
	}

	return generateAndSignJWT()
}

func generateAndSignJWT() error {
	rsaKey, err := LoadKey(keyFilePath)
	if err != nil {
		return err
	}

	key := jose.SigningKey{
		Algorithm: jose.SignatureAlgorithm(keyAlgorythm),
		Key:       rsaKey,
	}

	opts := &jose.SignerOptions{}
	opts.WithType("JWT")

	signer, err := jose.NewSigner(key, opts)
	if err != nil {
		return err
	}

	builder := jwt.Signed(signer).Claims(jwt.Claims{
		Issuer:   "webgame-cloud-api-dev",
		Subject:  fmt.Sprintf("%d", claimUserID),
		Audience: []string{jwtAudience},
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Expiry:   jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
	}).Claims(map[string]any{
		"userEmail": claimUserEmail,
	})

	token, err := builder.CompactSerialize()
	if err != nil {
		return err
	}

	fmt.Printf("Token: %s\n", token)

	return nil
}

func LoadKey(keyFilePath string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(keyData)
	if pemBlock == nil {
		return nil, fmt.Errorf("not valid private PEM key file")
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPrivKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key is not *rsa.PrivateKey")
	}
	return rsaPrivKey, nil
}
