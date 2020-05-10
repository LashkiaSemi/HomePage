package authentication

import (
	"homepage/pkg/configs"
	"net/http"
)

// CheckIsLogin 認証部分のチェック
func CheckIsLogin(r *http.Request) bool {
	_, err := r.Cookie(configs.CookieName)
	return (err == nil)
}
