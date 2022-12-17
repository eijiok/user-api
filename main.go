package main

import (
	"github.com/eijiok/user-api/server"
	"log"
)

func main() {
	err := server.InitServer()
	if err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
