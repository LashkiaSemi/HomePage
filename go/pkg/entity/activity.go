package entity

import (
	"homepage/pkg/configs"
	"time"
)

// Activity 活動内容
type Activity struct {
	ID          int
	Activity    string
	ShowDate    string
	LastDate    string
	Annotation  string
	IsImportant int
	CreatedAt   string
	UpdatedAt   string
}

func (a *Activity) Create(activity, showDate, lastDate, annotation string, isImportant int) {
	a.Activity = activity
	a.ShowDate = showDate
	a.LastDate = lastDate
	a.Annotation = annotation
	a.IsImportant = isImportant
	a.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	a.UpdatedAt = a.CreatedAt
}

func (a Activity) Update(activity, showDate, lastDate, annotation string, isImportant int) *Activity {
	a.Activity = activity
	a.ShowDate = showDate
	a.LastDate = lastDate
	a.Annotation = annotation
	a.IsImportant = isImportant
	a.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &a
}
