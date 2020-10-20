package entity

import (
	"homepage/pkg/configs"
	"time"
)

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

func (l *Lecture) Create(title, file, comment string, activation int, author *User) {
	l.Title = title
	l.File = file
	l.Comment = comment
	l.Activation = activation
	l.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	l.UpdatedAt = l.CreatedAt
	l.Author = author
}

func (l Lecture) Update(title, file, comment string, activation int, author *User) *Lecture {
	l.Title = title
	l.File = file
	l.Comment = comment
	l.Activation = activation
	l.Author = author
	return &l
}
