package utils

import (
	"fmt"
	"time"
)


// FormatDate takes a given date, localizes it, formats it in YYY-MM-DD
// form and returns the formatted date as a string
func FormatDate(date time.Time) string {
	// Use local time to give location-specific details
	return date.Local().Format("2006-01-02")
}

// PrintDate takes a given date and renders the time in a
// user friendly form to be printed as part of the UI
func PrintDate(date time.Time) string {
	localized := date.Local()
	weekday := localized.Weekday()
	month := localized.Month()
	day := localized.Day()

	return weekday.String() + ", " + month.String() + " " + fmt.Sprint(day)
}
