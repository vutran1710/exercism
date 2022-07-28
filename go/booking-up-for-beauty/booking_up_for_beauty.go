package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	var parsedTime, _ = time.Parse(layout, date)
	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	current, _ := time.Parse("January 2, 2006 15:04:05", date)
	return now.After(current)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointment, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := appointment.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	prefix := "You have an appointment on "
	timeStr, _ := time.Parse(layout, date)
	return prefix + fmt.Sprintf("%s", timeStr.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	date := fmt.Sprintf("%d-09-15", currentYear)
	layout := "2006-01-02"
	anni, _ := time.Parse(layout, date)
	return anni
}
