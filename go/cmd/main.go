package main

import (
	"homepage/pkg/infrastructure/server"
)

func main() {
	port := "8080"

	// TODO: connection db

	// TODO: make handler

	// make server
	serv := server.NewServer(port)

	// listen
	serv.Serve()

}
