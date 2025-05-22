package utils

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func GetDateStartAndEndOfMonth(month time.Month, year int) (time.Time, time.Time, error) {
	if year < 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("year %d is not valid", year)
	}

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().Location())

	end := start.AddDate(0, 1, -1)

	zap.S().Infof("date start, end of month: %#v, %#v\n", start, end)

	return start, end, nil
}
