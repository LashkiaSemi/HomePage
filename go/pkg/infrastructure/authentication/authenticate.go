package authentication

import (
	"golang.org/x/crypto/bcrypt"
	"homepage/pkg/usecase/interactor"
)

type authHandler struct{}

func NewAuthHandler() interactor.AuthHandler {
	return &authHandler{}
}

// PasswordHash パスワードのハッシュ
func (ah *authHandler) PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// PasswordVerify パスワードの認証
func (ah *authHandler) PasswordVerify(hash, pw string) error {
	// 認証に失敗した場合は error
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
