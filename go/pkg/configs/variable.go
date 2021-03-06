package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 環境変数から設定する変数たち
var (
	DBUser     string
	DBPassword string
	DBProtocol string
	DBTarget   string
	DBName     string

	JWTSecret string
)

// init 環境変数から読み込みが必要なものを初期化
func init() {
	err := godotenv.Load(".env")
	if err != nil && os.Getenv("ENV") != "test" {
		log.Fatal("[error] failed to find dotenv file")
	}

	// db関連
	DBUser = os.Getenv("MARIA_USER")
	DBPassword = os.Getenv("MARIA_PASSWORD")
	DBProtocol = os.Getenv("MARIA_PROTOCOL")
	DBTarget = os.Getenv("MARIA_TARGET")
	DBName = os.Getenv("MARIA_DB")

	// jwt
	JWTSecret = os.Getenv("JWT_SECRET")
}
