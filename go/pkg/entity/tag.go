package entity

import (
	"homepage/pkg/configs"
	"time"
)

// Tag タグ
type Tag struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
}

func (t *Tag) Create(name string) {
	t.Name = name
	t.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	t.UpdatedAt = t.CreatedAt
}

func (t Tag) Update(name string) *Tag {
	t.Name = name
	t.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &t
}
