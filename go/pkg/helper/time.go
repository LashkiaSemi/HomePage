package helper

import (
	"time"
)

const (
	dateTimeFormat = "2006/01/02 15:04:05"
)

// Now is get now
func Now() time.Time {
	return time.Now()
}

// FormattedDateTimeNow get current time of format(YYYY/MM/DD hh:mm:ss)
func FormattedDateTimeNow() string {
	return time.Now().Format(dateTimeFormat)
}
