package auth

import (
	"errors"
	"homepage/pkg/configs"
	"homepage/pkg/usecase/interactor"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
		return errors.New("failed to check admin session")
	}
	return nil
}

// GetStudentIDFromCookie cookieから学籍番号の取得
func GetStudentIDFromCookie(r *http.Request) string {
	cookie, err := r.Cookie(configs.CookieName)
	if err != nil {
		log.Println("Cookie: ", err)
		return ""
	}
	tokenString := cookie.Value

	// jwtの検証
	token, err := VerifyToken(tokenString)
	if err != nil {
		log.Println("failed to verify token: ", err)
	}
	return GetStudentIDFromJWT(token)
}

type verifyHandler struct{}

func NewVerifyHandler() interactor.VerifyHandler {
	return &verifyHandler{}
}

// PasswordHash パスワードのハッシュ
func (v *verifyHandler) PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// PasswordVerify パスワードの認証
func (v *verifyHandler) PasswordVerify(hash, pw string) error {
	// 認証に失敗した場合は error
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
