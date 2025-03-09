package main

import (
	"binvault/pkg/api"
	"binvault/pkg/api/helpers"
	"binvault/pkg/auth"
	"binvault/pkg/compression"
	"binvault/pkg/database"
	"binvault/pkg/tasks"
	"log"
	"runtime"
)

var workers = runtime.NumCPU()

func main() {

	helpers.Init()

	auth := auth.GetAuth()
	log.Println("Auth enabled", auth.Enabled)

	compression.Init()

	tasks.Run(workers)

	database.Init()
	api.StartServer()

}
