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

// func (l *Lecture) Create(title, file, comment string, activation int, author *User) {
// 	l.Title = title
// 	l.File = file
// 	l.Comment = comment
// 	l.Activation = activation
// 	l.CreatedAt = time.Now().Format(configs.DateTimeFormat)
// 	l.UpdatedAt = l.CreatedAt
// 	l.Author = author
// }

func (l Lecture) Update(title, comment string, activation int) *Lecture {
	res := l
	res.Title = title
	res.Comment = comment
	res.Activation = activation
	return &res
}
