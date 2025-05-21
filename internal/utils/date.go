package utils

import (
	"fmt"
	"time"
)

func GetDateStartAndEndOfMonth(month time.Month, year int) (time.Time, time.Time, error) {
	if year < 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("year %d is not valid", year)
	}

	start := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)

	end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)

	return start, end, nil
}
