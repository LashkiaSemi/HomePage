//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

// VerifyHandler パスワード検証を行うハンドラ
type VerifyHandler interface {
	// PasswordHash 生パスをハッシュする
	PasswordHash(pw string) (string, error)
	// PasswordVerify パスワードの検証をする
	PasswordVerify(hash, pw string) error
}
