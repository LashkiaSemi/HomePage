package auth

import (
	"errors"
	"homepage/pkg/configs"
	"log"
	"net/http"
)

// adminSessions studentID: tokenString?
var adminSessions = make(map[string]string)

// SetAdminSession adminログインしたときに、サーバに保存しておく
func SetAdminSession(studentID, token string) {
	adminSessions[studentID] = token
}

// CheckIsAdminSession 学籍番号からadminログインしたときのJWTを取得
func CheckIsAdminSession(studentID, token string) error {
	if adminSessions[studentID] != token {
		return errors.New("unmatch current token and admin-session token")
	}
	return nil
}

// GetStudentIDFromCookie cookieから学籍番号の取得
func GetStudentIDFromCookie(r *http.Request) string {
	cookie, err := r.Cookie(configs.CookieName)
	if err != nil {
		// log.Printf("failed to get cookie: %v", err)
		return ""
	}
	tokenString := cookie.Value

	// jwtの検証
	token, err := VerifyToken(tokenString)
	if err != nil {
		log.Printf("failed to verify token: %v", err)
		return ""
	}
	return GetStudentIDFromJWT(token)
}
