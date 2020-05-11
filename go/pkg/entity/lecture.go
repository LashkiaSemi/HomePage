package entity

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

func (l Lecture) Update(title, comment string, activation int) *Lecture {
	res := l
	res.Title = title
	res.Comment = comment
	res.Activation = activation
	return &res
}
