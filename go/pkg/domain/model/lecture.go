package model

// Lecture レクチャーのモデル
type Lecture struct {
	ID         int
	Author     *User
	Title      string
	File       string
	Comment    string
	Activation int
	CreatedAt  string
	UpdatedAt  string
}
