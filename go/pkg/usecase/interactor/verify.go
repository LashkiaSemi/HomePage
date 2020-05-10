package interactor

type VerifyHandler interface {
	PasswordHash(pw string) (string, error)
	PasswordVerify(hash, pw string) error
}
