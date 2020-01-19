package main

import (
	"homepage/conf"
	"homepage/pkg/infrastructure/authentication"
	"homepage/pkg/infrastructure/datastore"
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server"
	"homepage/pkg/infrastructure/server/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load config
	servConf := conf.LoadServerConfig()

	// connection db
	sh := datastore.NewSQLHandler()

	// create server
	serv := server.NewServer(servConf["host"], servConf["port"])

	// create app handler
	au := authentication.NewAuthHandler()
	ah := handler.NewAppHandler(sh, au)

	// routing
	router.SettingRouter(serv, ah)

	// server start
	serv.Serve()

}
