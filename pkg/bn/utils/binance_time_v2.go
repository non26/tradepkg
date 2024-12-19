package utils

import (
	"fmt"
	"time"
)

func GetBinanceTimestamp() int64 {
	return time.Now().Unix() * 1000
}

func GetSpecificBnTimestamp(t *time.Time) int64 {
	return t.Unix() * 1000
}

type binanceTime struct {
	start_date_time time.Time
	execute_at      time.Time
}

func NewBinanceTime(execution_time time.Time) *binanceTime {
	b := &binanceTime{}
	b.setBinanceStartDateTime(execution_time)
	b.SetExecutionTime(execution_time)
	return b
}

// t is server time
func (b *binanceTime) setBinanceStartDateTime(t time.Time) time.Time {
	utc_time := b.convertToUTC(t)
	bn_start_time := time.Date(utc_time.Year(), utc_time.Month(), utc_time.Day(), 0, 0, 0, 0, time.UTC)
	return bn_start_time
}

func (b *binanceTime) SetExecutionTime(t time.Time) {
	b.start_date_time = b.setBinanceStartDateTime(t)
	b.execute_at = t
}

func (b *binanceTime) convertToUTC(t time.Time) time.Time {
	return t.In(time.UTC)
}

func (b *binanceTime) GetBnTimeStartHourAndEndHour(period int) (time.Time, time.Time, error) {
	remainder := 24 % period
	if remainder != 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("period must be a divisor of 24")
	}

	period_count := (24 / period) - 1 // 24 is the number of hours in a day

	for i := 0; i <= period_count; i++ {
		start_hour_time := b.startPeriod(i, period, Hour)
		end_hour_time := b.endPeriod(start_hour_time, period, Hour)
		_, _, ok := isBetween(b.execute_at, start_hour_time, end_hour_time)
		if ok {
			return start_hour_time, end_hour_time, nil
		}
	}
	return time.Time{}, time.Time{}, fmt.Errorf("no start and end of period found")
}

func (b *binanceTime) GetPreviousBnTimeStartHourAndEndHour(period int) (time.Time, time.Time, error) {
	start_time, end_time, err := b.GetBnTimeStartHourAndEndHour(period)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return start_time.Add(time.Duration(-1*period) * time.Hour), end_time.Add(time.Duration(-1*period) * time.Hour), nil
}

func (b *binanceTime) GetBnTimeStartMinuteAndEndMinute(period int) (time.Time, time.Time, error) {
	remainder := 1440 % period
	if remainder != 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("period must be a divisor of 1440(24*60)")
	}

	period_count := (1440 / period) - 1 // 1440 is the number of minutes in a day

	for i := 0; i <= period_count; i++ {
		start_minute_time := b.startPeriod(i, period, Minute)
		end_minute_time := b.endPeriod(start_minute_time, period, Minute)
		_, _, ok := isBetween(b.execute_at, start_minute_time, end_minute_time)
		if ok {
			return start_minute_time, end_minute_time, nil
		}
	}
	return time.Time{}, time.Time{}, fmt.Errorf("no start and end of period found")
}

func (b *binanceTime) GetPreviousBnTimeStartMinuteAndEndMinute(period int) (time.Time, time.Time, error) {
	start_time, end_time, err := b.GetBnTimeStartMinuteAndEndMinute(period)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return start_time.Add(time.Duration(-1*period) * time.Minute), end_time.Add(time.Duration(-1*period) * time.Minute), nil
}

func (b *binanceTime) GetBnTimeStartDayAndEndDay(period int) (time.Time, time.Time, error) {
	start_day_time := b.startPeriod(0, period, Day)
	end_day_time := b.endPeriod(start_day_time, period, Day)
	return start_day_time, end_day_time, nil
}

func (b *binanceTime) GetPreviousBnTimeStartDayAndEndDay(period int) (time.Time, time.Time, error) {
	start_time, end_time, err := b.GetBnTimeStartDayAndEndDay(period)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	hour_count := period * 24
	start_time = start_time.Add(time.Duration(-hour_count) * time.Hour)
	end_time = end_time.Add(time.Duration(-hour_count) * time.Hour)
	return start_time, end_time, nil
}

func (b *binanceTime) startPeriod(round int, period int, unit time.Duration) time.Time {
	switch unit {
	case Minute:
		return b.start_date_time.Add(time.Minute * time.Duration(round*period))
	case Hour:
		return b.start_date_time.Add(time.Hour * time.Duration(round*period))
	case Day:
		return b.start_date_time
	default:
		return time.Time{}
	}
}

func (b *binanceTime) endPeriod(start_period time.Time, period int, unit time.Duration) time.Time {
	switch unit {
	case Minute:
		return start_period.Add(time.Minute * time.Duration(period-1)).Add(59 * time.Second)
	case Hour:
		return start_period.Add(time.Hour * time.Duration(period-1)).Add(59 * time.Minute).Add(59 * time.Second)
	case Day:
		return start_period.Add(time.Hour * time.Duration(period*24-1)).Add(59 * time.Minute).Add(59 * time.Second)
	default:
		return time.Time{}
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

// (24-1)*2 + 1 = 24(2) - 1(2) + 1 = 48-2+1 = 48-1 => 2(24) -1 = period * 24 - 1
