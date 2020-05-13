package entity

import (
	"homepage/pkg/configs"
	"time"
)

// Activity 活動内容
type Activity struct {
	ID        int
	Activity  string
	Date      string
	CreatedAt string
	UpdatedAt string
}

func (a *Activity) Create(activity, date string) {
	a.Activity = activity
	a.Date = date
	a.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	a.UpdatedAt = a.CreatedAt
}

func (a Activity) Update(activity, date string) *Activity {
	a.Activity = activity
	a.Date = date
	a.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &a
}
