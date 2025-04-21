package jwtgen

import (
	"binvault/pkg/env"
	"binvault/pkg/services/auth"
	"fmt"
	"log"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/spf13/cobra"
)

var JWTGen = &cobra.Command{
	Use:   "jwt",
	Short: "Generate JWT",
	RunE:  RunJWTGen,
}

var keyAlgorythm string
var claimUserID int64

func init() {
	flags := JWTGen.Flags()
	flags.StringVar(&keyAlgorythm, "alg", "RS256", "Hash algorithm")
	flags.Int64Var(&claimUserID, "id", 0, "Generate JWT for this user")
}

func RunJWTGen(cmd *cobra.Command, args []string) error {
	if claimUserID == 0 {
		return fmt.Errorf("plase set id")
	}

	return generateAndSignJWT()
}

func generateAndSignJWT() error {

	var privateKey = auth.LoadRSAPrivateKey()
	pem := auth.LoadPEM()

	if privateKey == nil && pem != nil {
		privateKey = pem.PrivateKey
	}

	if privateKey == nil && pem == nil {
		auth.GeneratePEM(env.GetVars().DATA_PATH)
		privateKey = auth.LoadPEM().PrivateKey
	}

	if privateKey == nil {
		return fmt.Errorf("failed to load private key")
	}

	signingKey := jose.SigningKey{
		Algorithm: jose.SignatureAlgorithm(keyAlgorythm),
		Key:       privateKey,
	}

	opts := &jose.SignerOptions{}
	opts.WithType("JWT")

	signer, err := jose.NewSigner(signingKey, opts)
	if err != nil {
		return err
	}

	builder := jwt.Signed(signer).Claims(jwt.Claims{
		Issuer:   "binvault",
		Subject:  fmt.Sprintf("%d", claimUserID),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Expiry:   jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
	}).Claims(map[string]any{
		"id": claimUserID,
	})

	token, err := builder.Serialize()
	if err != nil {
		return err
	}

	log.Printf("Token: %s\n", token)

	return nil
}
