package interactor

// VerifyHandler パスワード検証を行うハンドラ
type VerifyHandler interface {
	PasswordHash(pw string) (string, error)
	PasswordVerify(hash, pw string) error
}
