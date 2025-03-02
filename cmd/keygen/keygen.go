package keygen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/spf13/cobra"
)

var KeyGen = &cobra.Command{
	Use:   "keygen",
	Short: "Generate development key.",
	RunE:  RunKeyGen,
}

var keyFile string
var keyBits int

func init() {
	flags := KeyGen.Flags()

	flags.StringVar(&keyFile, "key-file", ".dev/key.pem", "Generated key file.")
	flags.IntVar(&keyBits, "bits", 4096, "Key bit size.")
}

func RunKeyGen(cmd *cobra.Command, args []string) error {
	f, err := os.Create(keyFile)
	if err != nil {
		return err
	}
	defer f.Close()

	key, err := rsa.GenerateKey(rand.Reader, keyBits)
	if err != nil {
		return err
	}

	err = pem.Encode(f, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	return err
}
