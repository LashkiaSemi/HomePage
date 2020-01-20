package interactor

// AuthHandler パスワードのハッシュとか検証。実装はinfraで行なっています。
type AuthHandler interface {
	PasswordHash(pw string) (string, error)
	PasswordVerify(hash, pw string) error
}
