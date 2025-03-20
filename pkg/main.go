package main

import (
	"binvault/pkg/api"
	"binvault/pkg/api/helpers"
	"binvault/pkg/database"
	"binvault/pkg/processors"
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
	log.Println("Authorization enabled:", auth.Enabled)

	processors.Init()

	database.Init()
	tasks.Run(workers)
	log.Println("===================================")
	log.Println("===                             ===")
	log.Println("===       BINVAULT STARTED      ===")
	log.Println("===                             ===")
	log.Println("===================================")
	api.StartServer()

}
