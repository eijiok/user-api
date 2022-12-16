package main

import (
	"github.com/eijiok/user-api/server"
	"log"
)

func main() {
	err := server.InitServer("api", "8080")
	if err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
