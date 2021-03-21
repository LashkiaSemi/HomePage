package entity

import (
	"homepage/pkg/helper"
)

// Research 卒業研究
type Research struct {
	ID         int
	Title      string
	Author     string
	File       string
	Comment    string
	Activation int
	CreatedAt  string
	UpdatedAt  string
}

func NewResearch(title, author, file, comment string, activation int) *Research {
	now := helper.FormattedDateTimeNow()
	return &Research{
		Title:      title,
		Author:     author,
		File:       file,
		Comment:    comment,
		Activation: activation,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (r Research) Update(title, author, file, comment string, activation int) *Research {
	r.Title = title
	r.Author = author
	r.File = file
	r.Comment = comment
	r.Activation = activation
	r.UpdatedAt = helper.FormattedDateTimeNow()
	return &r
}
