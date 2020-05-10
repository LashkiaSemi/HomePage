package main

import (
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server"
)

func main() {
	port := "8080"

	// TODO: connection db

	// TODO: make handler
	ah := handler.NewAppHandler()

	// make server
	serv := server.NewServer(port, ah)

	// listen
	serv.Serve()

}
