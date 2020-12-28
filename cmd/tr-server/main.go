package main

import (
	"log"
	"net/http"

	api "github.com/acrochet95/transport-rennes-be/internal/transport-rennes-api"
)

func main() {
	api.InitializeServer()

	router := api.InitializeRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
