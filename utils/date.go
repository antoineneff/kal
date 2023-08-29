package utils

import "time"

func MonthDetails(year int, month time.Month, location *time.Location) (int, int) {
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)
	endOfMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, location)

	weekday := startOfMonth.Weekday()
	numberOfDays := endOfMonth.Sub(startOfMonth).Hours() / 24

	return int(weekday), int(numberOfDays)
}
