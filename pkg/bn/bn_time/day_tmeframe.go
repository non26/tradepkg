package bntime

import (
	"strconv"
	"time"
)

func GetBinanceStartAndEndTimeInDayTimeFrame(number string) (start time.Time, end time.Time) {
	// numberInt, err := strconv.Atoi(number)
	// if err != nil {
	// 	return time.Time{}, time.Time{}
	// }

	targetTime := time.Now().UTC()

	start = time.Date(targetTime.Year(), targetTime.Month(), targetTime.Day(), 0, 0, 0, 0, time.UTC)
	end = time.Date(targetTime.Year(), targetTime.Month(), targetTime.Day(), 23, 59, 59, 59, time.UTC)

	return start, end
}

func GetBinancePreviousStartAndEndTimeInDayTimeFrame(
	start time.Time,
	end time.Time,
	number string) (previousStart time.Time, previousEnd time.Time) {
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		return time.Time{}, time.Time{}
	}

	previousStart = start.Add(-time.Duration(numberInt) * 24 * time.Hour)
	previousEnd = end.Add(-time.Duration(numberInt) * 24 * time.Hour)

	return previousStart, previousEnd
}
