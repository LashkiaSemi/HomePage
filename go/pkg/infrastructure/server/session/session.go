package session

import "github.com/gorilla/sessions"

// Store セッションを保存する場所
var Store = sessions.NewCookieStore([]byte("hoge"))
