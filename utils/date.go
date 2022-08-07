package utils

import (
	"time"
)

type Date string

// FormatDate takes a given date, localizes it, formats it in YYY-MM-DD
// form and returns the formatted date as a string
func FormatDate(date time.Time) string {
	// Use local time to give location-specific details
	return date.Local().Format("2006-01-02")
}
