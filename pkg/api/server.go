package api

import (
	"log"
	"net/http"
)

func StartServer() {
	router := initRouter()
	var err = http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
