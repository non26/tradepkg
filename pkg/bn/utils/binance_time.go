package utils

import (
	"time"

	"github.com/non26/tradepkg/pkg/bn/thaitime"
)

func GetBinanceTimestamp() int64 {
	return time.Now().Unix() * 1000
}

func GetSpecificBnTimestamp(t *time.Time) int64 {
	return t.Unix() * 1000
}

// _t is time.Now()
// TODO: period can be minute, hour, day, week
func GetUnixBnStartAndEndOfPeriodHoursIntraday(_t time.Time, thaitime *thaitime.ThaiTime, period int, unit time.Duration) (*time.Time, *time.Time) {
	for i := 0; i < 24; i += period {
		start := time.Date(_t.Year(), _t.Month(), _t.Day(), i, 0, 0, 0, thaitime.Location)
		end := endOfPeriod(start, period, unit)
		_, _, ok := isBetween(_t, start, end)
		if ok {
			return &start, &end
		}
	}
	return nil, nil
}

// _t is time.Now()
func GetPreviousUnixBnStartAndEndOfPeriodHours(_t time.Time, thaitime *thaitime.ThaiTime, period int, unit time.Duration) (*time.Time, *time.Time) {
	_070000, _, ok := IsFirstPeriodOfDay(_t, thaitime, period, unit)
	if ok {
		start_period := _070000.Add(time.Duration(1-period) * unit)
		end_period := endOfPeriod(start_period, period, unit)
		return &start_period, &end_period
	}
	start_period, _ := GetUnixBnStartAndEndOfPeriodHoursIntraday(_t, thaitime, period, unit)
	previous_start_period := start_period.Add(time.Duration(1-period) * unit)
	previous_end_period := endOfPeriod(previous_start_period, period, unit)
	return &previous_start_period, &previous_end_period
}

func endOfPeriod(t time.Time, period int, unit time.Duration) time.Time {
	switch unit {
	case time.Duration(1) * time.Minute:
		return t.Add(time.Duration(period-1) * time.Minute).Add(59 * time.Second).Add(999 * time.Millisecond)
	case time.Duration(1) * time.Hour:
		return t.Add(time.Duration(period-1) * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	case time.Duration(1) * time.Hour * 24:
		return t.Add(time.Duration(period-1) * time.Hour * 24).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	case time.Duration(1) * time.Hour * 24 * 7:
		return t.Add(time.Duration(period-1) * time.Hour * 24 * 7).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	default:
		return t
	}
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

// _t is time.Now(), period is in format of  1
func IsFirstPeriodOfDay(_t time.Time, thaitime *thaitime.ThaiTime, period int, unit time.Duration) (time.Time, time.Time, bool) {
	_070000 := time.Date(_t.Year(), _t.Month(), _t.Day(), 7, 0, 0, 0, thaitime.Location)
	endOf_070000 := endOfPeriod(_070000, period, unit)
	_, _, ok := isBetween(_t, _070000, endOf_070000)
	return _070000, endOf_070000, ok
}
