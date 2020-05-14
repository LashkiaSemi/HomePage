package configs

import "time"

// applicationのconfig
const (
	AppHost = "localhost"
	AppPort = "8080"
)

// db config
const (
	DBDriver   = "mysql"
	DBUser     = "root"
	DBPassword = "password"
	DBProtocol = "tcp"
	DBTarget   = "localhost:3307"
	DBName     = "homepage"
)

// file upload先
const (
	// TODO: 環境変数にした方がいい。絶対パスなら任意の場所に突っ込める

	// SaveLectureFileDir レクチャー資料のアップロード先
	SaveLectureFileDir = "./public/lectures"

	// SaveResearchFileDir 卒業研究資料のアップロード先
	SaveResearchFileDir = "./public/researches"
)

// Session系
const (
	// JWTSecret jwtのシークレットキー
	// TODO: 環境変数とか...
	JWTSecret = "secret"

	// JWTStudentIDClaim 学籍番号のクレームのキー
	JWTStudentIDClaim = "student_id"

	// JWTExpire jwtの有効期限
	// とりあえず1日
	JWTExpire = time.Hour * 24

	// CookieName cookieのキー
	CookieName = "_l_semi_homepage_session"
)

// ほか
const (
	// DateTimeFormat 日付のフォーマット
	DateTimeFormat = "2006-01-02 15:4:5"
	// デフォルトのモード[release, admin]とかから？
	DefaultMode = "release"
)

// ModePtr 実行モード
var ModePtr = "release"
