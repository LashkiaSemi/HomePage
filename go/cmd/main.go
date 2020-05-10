package main

import (
	"homepage/pkg/infrastructure/database"
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server"
)

func main() {
	port := "8080"

	// connection db
	sh := database.NewSQLHandler()

	// create handler
	ah := handler.NewAppHandler(sh)

	// make server
	serv := server.NewServer(port, ah)

	// listen
	serv.Serve()

}
