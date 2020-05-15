package entity

import (
	"homepage/pkg/configs"
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

func (s *Society) Create(title, author, society, award, date string) {
	s.Title = title
	s.Author = author
	s.Society = society
	s.Award = award
	s.Date = date
	s.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	s.UpdatedAt = s.CreatedAt
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
