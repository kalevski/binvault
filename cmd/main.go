package main

import (
	"binvault/cmd/jwtgen"
	"binvault/cmd/keygen"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var WGCliCommand = &cobra.Command{
	Use:   "wgcli",
	Short: "WebGame cloud command line untility.",
}

func main() {
	WGCliCommand.AddCommand(jwtgen.JWTGen)
	WGCliCommand.AddCommand(keygen.KeyGen)

	if err := WGCliCommand.Execute(); err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}
