package session

import (
	"github.com/gorilla/sessions"
	"homepage/conf"
)

// Store sessionの保存先。cookie
var Store *sessions.CookieStore

func init() {
	Store = sessions.NewCookieStore([]byte("hoge"))
	Store.Options.MaxAge = conf.CookieMaxAge
	// Store.Options.Secure = true
	Store.Options.HttpOnly = true
}
