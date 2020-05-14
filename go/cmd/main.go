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
		// TODO: ひでえ英語だ
		log.Printf("[info] mode of '%v' is nothing. set mode '%v'.", *mode, configs.DefaultMode)
	} else {
		configs.ModePtr = *mode
	}
	log.Printf("[info] running mode: %v", configs.ModePtr)
}

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
