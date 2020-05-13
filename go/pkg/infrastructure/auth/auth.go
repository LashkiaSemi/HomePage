package auth

import (
	"homepage/pkg/usecase/interactor"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type verifyHandler struct{}

// NewVerifyHandler パスワード周りを担当するハンドラを作成
func NewVerifyHandler() interactor.VerifyHandler {
	return &verifyHandler{}
}

// PasswordHash パスワードのハッシュ
func (v *verifyHandler) PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		err = errors.Wrap(err, "failed to generate hash")
		return "", err
	}
	return string(hash), nil
}

// PasswordVerify パスワードの認証
func (v *verifyHandler) PasswordVerify(hash, pw string) error {
	// 認証に失敗した場合は error
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
