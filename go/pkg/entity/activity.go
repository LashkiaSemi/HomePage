package entity

import (
	"homepage/pkg/configs"
	"time"
)

// Activity 活動内容
type Activity struct {
	ID        int
	Activity  string
	ShowDate  string
	LastDate  string
	CreatedAt string
	UpdatedAt string
}

func (a *Activity) Create(activity, showDate, lastDate string) {
	a.Activity = activity
	a.ShowDate = showDate
	a.LastDate = lastDate
	a.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	a.UpdatedAt = a.CreatedAt
}

func (a Activity) Update(activity, showDate, lastDate string) *Activity {
	a.Activity = activity
	a.ShowDate = showDate
	a.LastDate = lastDate
	a.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &a
}
