package entity

import (
	"homepage/pkg/configs"
	"homepage/pkg/helper"
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

func NewEquipment(name, comment string, stock, tagID int) *Equipment {
	now := helper.FormattedDateTimeNow()
	return &Equipment{
		Name:      name,
		Stock:     stock,
		Comment:   comment,
		CreatedAt: now,
		UpdatedAt: now,
		Tag: &Tag{
			ID: tagID,
		},
	}
}

func (e Equipment) Update(name, comment string, stock, tagID int) *Equipment {
	e.Name = name
	e.Comment = comment
	e.Stock = stock
	e.Tag.ID = tagID
	e.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &e
}
