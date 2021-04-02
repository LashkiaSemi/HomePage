package entity

import (
	"homepage/pkg/helper"
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

// TODO: int -> bool
func NewActivity(activity, showDate, date, annotation string, isImportant, isNotify int) *Activity {
	now := helper.FormattedDateTimeNow()
	return &Activity{
		Activity:    activity,
		ShowDate:    showDate,
		Date:        date,
		Annotation:  annotation,
		IsImportant: isImportant,
		IsNotify:    isNotify,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (a Activity) Update(activity, showDate, date, annotation string, isImportant, isNotify int) *Activity {
	a.Activity = activity
	a.ShowDate = showDate
	a.Date = date
	a.Annotation = annotation
	a.IsImportant = isImportant
	a.IsNotify = isNotify
	a.UpdatedAt = helper.FormattedDateTimeNow()
	return &a
}
