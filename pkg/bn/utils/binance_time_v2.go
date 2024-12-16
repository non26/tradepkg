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

func (b *binanceTime) startHourPeriod(round int, period int) time.Time {
	return b.start_date_time.Add(time.Hour * time.Duration(round*period))
}

func (b *binanceTime) endHourPeriod(start_hour_time time.Time, period int) time.Time {
	return start_hour_time.Add(time.Hour * time.Duration(period-1)).Add(59 * time.Minute).Add(59 * time.Second)
}

func (b *binanceTime) GetBnTimeStartHourAndEndHour(period int) (time.Time, time.Time, error) {
	remainder := 24 % period
	if remainder != 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("period must be a divisor of 24")
	}

	period_count := (24 / period) - 1

	for i := 0; i <= period_count; i++ {
		start_hour_time := b.startHourPeriod(i, period)
		end_hour_time := b.endHourPeriod(start_hour_time, period)
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

func (b *binanceTime) startMinutePeriod(round int, period int) time.Time {
	return b.start_date_time.Add(time.Minute * time.Duration(round*period))
}

func (b *binanceTime) endMinutePeriod(start_minute_time time.Time, period int) time.Time {
	return start_minute_time.Add(time.Minute * time.Duration(period-1)).Add(59 * time.Second)
}

func (b *binanceTime) GetBnTimeStartMinuteAndEndMinute(period int) (time.Time, time.Time, error) {
	remainder := 60 % period
	if remainder != 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("period must be a divisor of 60")
	}

	period_count := (60 / period) - 1

	for i := 0; i <= period_count; i++ {
		start_minute_time := b.startMinutePeriod(i, period)
		end_minute_time := b.endMinutePeriod(start_minute_time, period)
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

func isBetween(t time.Time, start time.Time, end time.Time) (int64, int64, bool) {
	startUnix := start.Unix()
	endUnix := end.Unix()
	tUnix := t.Unix()
	if startUnix <= tUnix && tUnix <= endUnix {
		return startUnix, endUnix, true
	}
	return 0, 0, false
}
