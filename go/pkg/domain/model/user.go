package model

// User ユーザ
type User struct {
	ID         int
	StudentID  string
	Name       string
	Password   string
	Role       string
	Department string
	Grade      int
	Comment    string
	CreatedAt  string
	UpdatedAt  string
}
