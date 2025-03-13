package main

import (
	"binvault/pkg/api"
	"binvault/pkg/api/helpers"
	"binvault/pkg/compression"
	"binvault/pkg/database"
	"binvault/pkg/services/auth"
	"binvault/pkg/services/filesystem"
	"binvault/pkg/tasks"
	"log"
	"runtime"
)

var workers = runtime.NumCPU()

func main() {

	filesystem.Init()
	helpers.Init()

	auth := auth.GetAuth()
	log.Println("Auth enabled", auth.Enabled)

	compression.Init()

	database.Init()
	tasks.Run(workers)
	log.Println("=== BINVAULT STARTED ===")
	api.StartServer()

}
