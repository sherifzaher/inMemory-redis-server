package main

import (
	"github.com/sherifzaher/inMemory-redis-server/pkg/server"
	"log"
)

var serverAddress = ":8080"

func main() {
	srv := server.New(serverAddress)
	if err := srv.Start(); err != nil {
		panic(err)
	}
	log.Printf("Server is running on localhost%s", serverAddress)
}
