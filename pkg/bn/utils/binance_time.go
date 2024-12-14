package utils

import (
	"time"

	"github.com/non26/tradepkg/pkg/bn/thaitime"
)

func GetBinanceTimestamp() int64 {
	return time.Now().Unix() * 1000
}

func GetUnixBinanceStartAndEndOfPeriodHours(_t time.Time, thaitime *thaitime.ThaiTime, period int) (int64, int64) {
	for i := 0; i < 24; i += period {
		start := time.Date(_t.Year(), _t.Month(), _t.Day(), i, 0, 0, 0, thaitime.Location)
		end := endOfHours(start, period)
		startUnix, endUnix, ok := isBetween(_t, start, end)
		if ok {
			return startUnix, endUnix
		}
	}
	return 0, 0
}

func endOfHours(t time.Time, unit int) time.Time {
	return t.Add(time.Duration(unit-1) * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}

func isBetween(t time.Time, start time.Time, end time.Time) (int64, int64, bool) {
	startUnix := start.Unix()
	endUnix := end.Unix()
	tUnix := t.Unix()
	if startUnix <= tUnix && tUnix <= endUnix {
		return startUnix, endUnix, true
	}
	return 0, 0, false
}
