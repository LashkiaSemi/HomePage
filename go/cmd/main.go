package main

import (
	"flag"
	"homepage/pkg/configs"
	"homepage/pkg/infrastructure/database"
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server"
	"log"
)

// ModePtr 実行モード。デバッグするときに使うやつ
var ModePtr *string

func init() {
	// logのカスタム
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Printf("[info] finish setting log")

	// run modeの設定
	ModePtr = flag.String("mode", configs.DefaultMode, "run mode. value=[admin, release]")
	flag.Parse()
	if *ModePtr != "release" && *ModePtr != "admin" {
		// TODO: ひでえ英語だ
		log.Printf("[info] mode of '%v' is nothing. set mode '%v'.", *ModePtr, configs.DefaultMode)
		*ModePtr = configs.DefaultMode
	}
	log.Printf("[info] running mode: %v", *ModePtr)

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
