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
	Date        string
	Annotation  string
	IsImportant int
	IsNotify    int
	CreatedAt   string
	UpdatedAt   string
}

func (a *Activity) Create(activity, showDate, date, annotation string, isImportant, isNotify int) {
	a.Activity = activity
	a.ShowDate = showDate
	a.Date = date
	a.Annotation = annotation
	a.IsImportant = isImportant
	a.IsNotify = isNotify
	a.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	a.UpdatedAt = a.CreatedAt
}

func (a Activity) Update(activity, showDate, date, annotation string, isImportant, isNotify int) *Activity {
	a.Activity = activity
	a.ShowDate = showDate
	a.Date = date
	a.Annotation = annotation
	a.IsImportant = isImportant
	a.IsNotify = isNotify
	a.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &a
}
