package main

import (
	"log"
)

func main() {
	httpServer, err := InitializeServer()
	if err != nil {
		log.Fatalf("failed to create http server: %s", err)
	}
	httpServer.Run()
}
