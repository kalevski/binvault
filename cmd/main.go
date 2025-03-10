package main

import (
	"binvault/cmd/jwtgen"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var WGCliCommand = &cobra.Command{
	Use:   "binvault",
	Short: "BinVault command line untility.",
}

func main() {
	WGCliCommand.AddCommand(jwtgen.JWTGen)

	if err := WGCliCommand.Execute(); err != nil {
		log.Println("error:", err.Error())
		os.Exit(1)
	}
}
