package bntime

import (
	"strconv"
	"time"
)

func GetBinanceStartAndEndTimeInHourTimeFrame(number string) (start time.Time, end time.Time) {
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		return time.Time{}, time.Time{}
	}
	if 24%numberInt != 0 {
		return time.Time{}, time.Time{}
	}
	startRound := 24 / numberInt

	targetTime := time.Now().UTC()
	targetHour := targetTime.Hour()

	initStartTime := 0
	initEndTime := 1 * numberInt

	for i := 0; i < startRound; i++ {
		if initStartTime <= targetHour && targetHour < initEndTime {
			start = time.Date(targetTime.Year(), targetTime.Month(), targetTime.Day(), initStartTime, 0, 0, 0, time.UTC)
			end = time.Date(targetTime.Year(), targetTime.Month(), targetTime.Day(), initEndTime, 59, 59, 59, time.UTC)
			break

		}
		initStartTime = initStartTime + numberInt
		initEndTime = initEndTime + numberInt
	}

	return start, end
}

func GetBinancePreviousStartAndEndTimeInHourTimeFrame(
	start time.Time,
	end time.Time,
	number string) (previousStart time.Time, previousEnd time.Time) {
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		return time.Time{}, time.Time{}
	}

	previousStart = start.Add(-time.Duration(numberInt) * time.Hour)
	previousEnd = end.Add(-time.Duration(numberInt+1) * time.Hour)

	return previousStart, previousEnd
}
