package auth

import (
	"homepage/pkg/configs"
	"homepage/pkg/usecase/interactor"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// CheckIsLogin 認証部分のチェック
func CheckIsLogin(r *http.Request) bool {
	_, err := r.Cookie(configs.CookieName)
	return (err == nil)
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
