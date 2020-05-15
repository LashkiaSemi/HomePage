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
		return errors.New("current token and admin-session token do not match")
	}
	return nil
}

// SetNewCookie 新しいCookieを作ってセット
func SetNewCookie(w http.ResponseWriter, value string) {
	cookie := &http.Cookie{
		Name:  configs.CookieName,
		Value: value,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}

// DeleteCookie Cookieを無効にする
func DeleteCookie(w http.ResponseWriter, cookie *http.Cookie) {
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
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
