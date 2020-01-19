package interactor

type AuthHandler interface {
	PasswordHash(pw string) (string, error)
	PasswordVerify(hash, pw string) error
}
