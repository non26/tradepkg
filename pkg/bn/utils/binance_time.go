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

func EndOfMinute(t time.Time, period int) time.Time {
	return t.Add(time.Duration(period-1) * time.Minute).Add(59 * time.Second)
}

func EndOfHour(t time.Time, period int) time.Time {
	return t.Add(time.Duration(period-1) * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}

func EndOfDay(t time.Time, period int) time.Time {
	return t.Add(time.Duration(period-1) * time.Hour * 24).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}

func EndOfWeek(t time.Time, period int) time.Time {
	return t.Add(time.Duration(period-1) * time.Hour * 24 * 7).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}

func EndOfMonth(t time.Time, period int) time.Time {
	// Move to the first day of the next month
	nextMonth := t.AddDate(0, 1, -t.Day()+1)
	// Subtract one day to get the last day of the current month
	lastDay := nextMonth.AddDate(0, 0, -1)
	return lastDay.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}

func bnNewDayForThaiTime(t time.Time, thaitime *thaitime.ThaiTime) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 7, 0, 0, 0, thaitime.Location)
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
	case Minute:
		return EndOfMinute(t, period)
	case Hour:
		return EndOfHour(t, period)
	case Day:
		return EndOfDay(t, period)
	case Week:
		return EndOfWeek(t, period)
	case Month:
		return EndOfMonth(t, period)
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
	_070000 := bnNewDayForThaiTime(_t, thaitime)
	endOf_070000 := endOfPeriod(_070000, period, unit)
	_, _, ok := isBetween(_t, _070000, endOf_070000)
	return _070000, endOf_070000, ok
}
