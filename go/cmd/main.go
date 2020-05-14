package main

import (
	"flag"
	"homepage/pkg/configs"
	"homepage/pkg/infrastructure/database"
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server"
	"log"
)

func init() {
	// logのカスタム
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("[info] finish setting log")

	// run modeの設定
	mode := flag.String("mode", configs.DefaultMode, "run mode. value=[admin, release]")
	flag.Parse()
	if *mode != "release" && *mode != "admin" {
		log.Printf("[info] the specified mode '%v' is invalid. set mode '%v'.", *mode, configs.DefaultMode)
	} else {
		configs.ModePtr = *mode
	}
	log.Printf("[info] running mode: %v", configs.ModePtr)
}

func main() {
	host := configs.AppHost
	port := configs.AppPort

	// connection db
	sh := database.NewSQLHandler()

	// create handler
	ah := handler.NewAppHandler(sh)

	// make server
	serv := server.NewServer(host, port, ah)

	// listen
	serv.Serve()

}
