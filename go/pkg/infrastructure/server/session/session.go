package session

import (
	"homepage/conf"
	"homepage/pkg/domain"

	"github.com/gorilla/sessions"
)

// Store sessionの保存先。cookie
var Store *sessions.CookieStore

func init() {
	Store = sessions.NewCookieStore([]byte("hoge"))
	Store.Options.MaxAge = conf.CookieMaxAge
	// Store.Options.Secure = true
	Store.Options.HttpOnly = true
}

// 使ってない...
func SetSessionData(sessData *domain.Session, sess *sessions.Session) {
	sess.Values["sessionID"] = sessData.SessionID
	sess.Values["userID"] = sessData.UserID
	sessData.SetSessionList()
}
