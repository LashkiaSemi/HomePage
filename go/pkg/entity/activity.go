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
	FirstDate string
	CreatedAt string
	UpdatedAt string
}

func (a *Activity) Create(activity, showDate, firstDate string) {
	a.Activity = activity
	a.ShowDate = showDate
	a.FirstDate = firstDate
	a.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	a.UpdatedAt = a.CreatedAt
}

func (a Activity) Update(activity, showDate, firstDate string) *Activity {
	a.Activity = activity
	a.ShowDate = showDate
	a.FirstDate = firstDate
	a.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &a
}
