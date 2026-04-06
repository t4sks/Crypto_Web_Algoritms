package main

import (
	"Crypto-Ciphers/internal/httpserver"
	"log"
)

func main() {
	server := httpserver.New()
	log.Fatal(server.ListenAndServe())
}
