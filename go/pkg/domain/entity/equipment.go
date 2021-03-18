package entity

import (
	"homepage/pkg/configs"
	"time"
)

// Equipment 備品
type Equipment struct {
	ID        int
	Name      string
	Stock     int
	Comment   string
	Tag       *Tag
	CreatedAt string
	UpdatedAt string
}

func (e *Equipment) Create(name, comment string, stock, tagID int) {
	tag := Tag{}
	tag.ID = tagID
	e.Name = name
	e.Stock = stock
	e.Comment = comment
	e.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	e.UpdatedAt = e.CreatedAt
	e.Tag = &tag
}

func (e Equipment) Update(name, comment string, stock, tagID int) *Equipment {
	e.Name = name
	e.Comment = comment
	e.Stock = stock
	e.Tag.ID = tagID
	e.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &e
}
