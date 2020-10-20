package configs

import (
	"os"
)

// 環境変数から設定する変数たち
var (
	DBUser     = DefaultDBUser
	DBPassword = DefaultDBPassword
	DBProtocol = DefaultDBProtocol
	DBTarget   = DefaultDBTarget
	DBName     = DefaultDBName

	JWTSecret = DefaultJWTSecret
)

// init 環境変数から読み込みが必要なものを初期化
func init() {
	// db関連
	if os.Getenv("MARIA_USER") != "" {
		DBUser = os.Getenv("MARIA_USER")
	}
	if os.Getenv("MARIA_PASSWORD") != "" {
		DBPassword = os.Getenv("MARIA_PASSWORD")
	}
	if os.Getenv("MARIA_PROTOCOL") != "" {
		DBProtocol = os.Getenv("MARIA_PROTOCOL")
	}
	if os.Getenv("MARIA_TARGET") != "" {
		DBTarget = os.Getenv("MARIA_TARGET")
	}
	if os.Getenv("MARIA_DB") != "" {
		DBName = os.Getenv("MARIA_DB")
	}

	// jwt
	if os.Getenv("JWT_SECRET") != "" {
		JWTSecret = os.Getenv("JWT_SECRET")
	}
}
