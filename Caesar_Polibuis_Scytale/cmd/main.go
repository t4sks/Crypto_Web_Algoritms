package main

import (
	"Polibuis_Scytale/internal/httpserver"
	"log"
)

func main() {
	server := httpserver.New()
	log.Fatal(server.ListenAndServe())
}
