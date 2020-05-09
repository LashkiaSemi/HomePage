package main

import (
	"homepage/conf"
	"homepage/pkg/domain/logger"
	"homepage/pkg/infrastructure/authentication"
	"homepage/pkg/infrastructure/datastore"
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server"
	"homepage/pkg/infrastructure/server/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// setup logfile
	err := logger.SetUpLogfile(conf.ServerLogfile)
	if err != nil {
		logger.Error("fail setup logfile")
		return
	}

	// connection db
	sh := datastore.NewSQLHandler()

	// create server
	serv := server.NewServer(conf.ServerHost, conf.ServerPort)

	// create app handler
	au := authentication.NewAuthHandler()
	ah := handler.NewAppHandler(sh, au)

	// routing
	router.SettingRouter(serv, ah)

	// server start
	serv.Serve()

	// close logfile
	logger.CloseLogfile()
}
