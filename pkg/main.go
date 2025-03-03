package main

import (
	"binvault/pkg/api"
	"binvault/pkg/auth"
	"binvault/pkg/clients/filesystem"
	"binvault/pkg/compression"
	"binvault/pkg/database"
	"log"
	"runtime"
)

var workers = runtime.NumCPU()

func main() {
	auth := auth.GetAuth()
	log.Println("Auth enabled", auth.Enabled)

	go filesystem.WatchFolder("temp")
	go compression.Init(workers)

	database.Init()
	api.StartServer()

}
