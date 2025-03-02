package main

import (
	"binvault/pkg/auth"
	"fmt"
	"runtime"
)

var workers = runtime.NumCPU()

func main() {
	// auth := cfg.GetAuth()
	// log.Println("Auth enabled", auth.Enabled)

	// go filesystem.WatchFolder("temp")
	// go compression.Init(workers)

	// database.Init()
	// api.StartServer()

	private := auth.LoadRSAPrivateKey()

	token, err := auth.GenerateJWT(private, map[string]interface{}{
		"sub":   "1234567890",
		"name":  "John Doe",
		"admin": true,
	})
	if err != nil {
		panic(err)
	}

	public := auth.LoadRSAPublicKey()
	claims, err := auth.ValidateJWT(public, token)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Claims: %+v\n", claims)

}
