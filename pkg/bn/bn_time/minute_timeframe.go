package bntime

import (
	"strconv"
	"time"
)

func GetBinanceStartAndEndTimeInMinuteTimeFrame(number string) (start time.Time, end time.Time) {
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		return time.Time{}, time.Time{}
	}
	if 60%numberInt != 0 {
		return time.Time{}, time.Time{}
	}
	startRound := 60 / numberInt

	targetTime := time.Now().UTC()
	targetMinute := targetTime.Minute()

	initStartTime := 0
	initEndTime := 1 * numberInt

	for i := 0; i < startRound; i++ {
		if initStartTime <= targetMinute && targetMinute < initEndTime {
			start = time.Date(targetTime.Year(), targetTime.Month(), targetTime.Day(), targetTime.Hour(), initStartTime, 0, 0, time.UTC)
			end = time.Date(targetTime.Year(), targetTime.Month(), targetTime.Day(), targetTime.Hour(), initStartTime, 59, 59, time.UTC)
			break
		}
		initStartTime = initStartTime + numberInt
		initEndTime = initEndTime + numberInt
	}

	return start, end
}
