package entity

import (
	"homepage/pkg/helper"
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

func NewLecture(title, file, comment string, activation int, author *User) *Lecture {
	now := helper.FormattedDateTimeNow()
	return &Lecture{
		Title:      title,
		Author:     author,
		File:       file,
		Comment:    comment,
		Activation: activation,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (l Lecture) Update(title, file, comment string, activation int, author *User) *Lecture {
	l.Title = title
	l.File = file
	l.Comment = comment
	l.Activation = activation
	l.Author = author
	l.UpdatedAt = helper.FormattedDateTimeNow()
	return &l
}
