package entity

import (
	"homepage/pkg/helper"
)

// Tag タグ
type Tag struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
}

func NewTag(name string) *Tag {
	now := helper.FormattedDateTimeNow()
	return &Tag{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t Tag) Update(name string) *Tag {
	t.Name = name
	t.UpdatedAt = helper.FormattedDateTimeNow()
	return &t
}
