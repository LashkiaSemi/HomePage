package entity

import (
	"homepage/pkg/configs"
	"homepage/pkg/helper"
	"time"
)

// Society 学会発表
type Society struct {
	ID        int
	Title     string
	Author    string
	Society   string
	Award     string
	Date      string
	CreatedAt string
	UpdatedAt string
}

func NewSociety(title, author, society, award, date string) *Society {
	now := helper.FormattedDateTimeNow()
	return &Society{
		Title:     title,
		Author:    author,
		Society:   society,
		Award:     award,
		Date:      date,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (s Society) Update(title, author, society, award, date string) *Society {
	s.Title = title
	s.Author = author
	s.Society = society
	s.Award = award
	s.Date = date
	s.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &s
}
