package configs

import (
	"time"
)

// applicationのconfig
const (
	AppHost = "0.0.0.0"
	AppPort = "8080"
)

// db config
const (
	DBDriver = "mysql"

	DefaultDBUser     = "root"
	DefaultDBPassword = "password"
	DefaultDBProtocol = "tcp"
	DefaultDBTarget   = "localhost:13306"
	DefaultDBName     = "homepage"
)

// file upload先
const (
	// 環境変数にした方がいい？
	// 絶対パスなら任意の場所に突っ込める
	PublicDir = "public/"
	StaticDir = "static/"

	// SaveLectureFileDir レクチャー資料のアップロード先
	SaveLectureFileDir = PublicDir + "lectures"

	// SaveResearchFileDir 卒業研究資料のアップロード先
	SaveResearchFileDir = PublicDir + "researches"
)

// Session系
const (
	// DefaultJWTSecret jwtのシークレットキー
	DefaultJWTSecret = "secret"

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
	DateTimeFormat = "2006/01/02 15:4:5"

	// DateForFileName ファイルアップロード時にファイル名かぶりを抑えるためのサフィックス用日付フォーマット
	DateForFileName = "2006-01-02-15-4-5"

	// デフォルトのモード[release, admin]とかから？
	DefaultMode = "release"
)

// ModePtr 実行モード
var ModePtr = "release"
