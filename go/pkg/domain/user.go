package domain

// User 一人分のデータ
type User struct {
	ID         int
	Name       string
	Password   string
	Role       string
	StudentID  string
	Department string
	Grade      int
	Comment    string
	CreatedAt  string
	UpdatedAt  string
}

// Users Userの配列
type Users []User
