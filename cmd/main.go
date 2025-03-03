package main

import (
	"binvault/cmd/jwtgen"
	"fmt"
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
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}
