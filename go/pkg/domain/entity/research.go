package entity

import (
	"homepage/pkg/configs"
	"time"
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

func (r *Research) Create(title, author, file, comment string, activation int) {
	r.Title = title
	r.Author = author
	r.File = file
	r.Comment = comment
	r.Activation = activation
	r.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	r.UpdatedAt = r.CreatedAt
}

func (r Research) Update(title, author, file, comment string, activation int) *Research {
	r.Title = title
	r.Author = author
	r.File = file
	r.Comment = comment
	r.Activation = activation
	r.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &r
}
