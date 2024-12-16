package utils

import (
	"testing"
	"time"

	"github.com/non26/tradepkg/pkg/bn/thaitime"
	"github.com/stretchr/testify/assert"
)

func TestBinanceTime_GetSpecificBnTimestamp(t *testing.T) {

	t.Run("GetSpecificBnTimestamp method", func(t *testing.T) {
		_t := time.Now()
		bn_timestamp := GetSpecificBnTimestamp(&_t)
		assert.Equal(t, bn_timestamp, _t.Unix()*1000)
	})
}

func TestBinanceTime_EndOfHour_1HoursPeriod(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	period := 1
	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 07:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 7, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time) // 07:00:00
		expected_end := EndOfHour(expected_start, period)    // 07:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 08:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 8, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(1 * time.Hour) // 08:00:00
		expected_end := EndOfHour(expected_start, period)                       // 08:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 09:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 9, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(2 * time.Hour) // 09:00:00
		expected_end := EndOfHour(expected_start, period)                       // 09:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 10:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 10, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(3 * time.Hour) // 10:00:00
		expected_end := EndOfHour(expected_start, period)                       // 10:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 11:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 11, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(4 * time.Hour) // 11:00:00
		expected_end := EndOfHour(expected_start, period)                       // 11:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 12:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 12, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(5 * time.Hour) // 12:00:00
		expected_end := EndOfHour(expected_start, period)                       // 12:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 13:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 13, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(6 * time.Hour) // 13:00:00
		expected_end := EndOfHour(expected_start, period)                       // 13:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 14:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 14, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(7 * time.Hour) // 14:00:00
		expected_end := EndOfHour(expected_start, period)                       // 14:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 15:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 15, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(8 * time.Hour) // 15:00:00
		expected_end := EndOfHour(expected_start, period)                       // 15:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 16:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 16, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(9 * time.Hour) // 16:00:00
		expected_end := EndOfHour(expected_start, period)                       // 16:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 17:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 17, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(10 * time.Hour) // 17:00:00
		expected_end := EndOfHour(expected_start, period)                        // 17:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 18:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 18, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(11 * time.Hour) // 18:00:00
		expected_end := EndOfHour(expected_start, period)                        // 18:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 19:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 19, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(12 * time.Hour) // 19:00:00
		expected_end := EndOfHour(expected_start, period)                        // 19:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 20:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 20, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(13 * time.Hour) // 20:00:00
		expected_end := EndOfHour(expected_start, period)                        // 20:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 21:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 21, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(14 * time.Hour) // 21:00:00
		expected_end := EndOfHour(expected_start, period)                        // 21:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 22:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 22, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(15 * time.Hour) // 22:00:00
		expected_end := EndOfHour(expected_start, period)                        // 22:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 23:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 23, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(16 * time.Hour) // 23:00:00
		expected_end := EndOfHour(expected_start, period)                        // 23:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 00:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 00, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 0, 0, 0, 0, thai_time.Location) // 0:00:00
		expected_end := EndOfHour(expected_start, period)                             // 0:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 01:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 1, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 1, 0, 0, 0, thai_time.Location) // 1:00:00
		expected_end := EndOfHour(expected_start, period)                             // 1:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 02:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 2, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 2, 0, 0, 0, thai_time.Location) // 2:00:00
		expected_end := EndOfHour(expected_start, period)                             // 2:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 03:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 3, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 3, 0, 0, 0, thai_time.Location) // 3:00:00
		expected_end := EndOfHour(expected_start, period)                             // 3:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 04:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 4, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 4, 0, 0, 0, thai_time.Location) // 4:00:00
		expected_end := EndOfHour(expected_start, period)                             // 4:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 05:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 5, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 5, 0, 0, 0, thai_time.Location) // 5:00:00
		expected_end := EndOfHour(expected_start, period)                             // 5:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 1 hours period: 06:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 6, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 6, 0, 0, 0, thai_time.Location) // 6:00:00
		expected_end := EndOfHour(expected_start, period)                             // 6:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTime_GetUnixBnStartAndEndOfPeriodHours_2HoursPeriod(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	period := 2
	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 08:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 8, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time) // 07:00:00
		expected_end := EndOfHour(expected_start, period)    // 08:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 10:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 10, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(2 * time.Hour) // 09:00:00
		expected_end := EndOfHour(expected_start, period)                       // 10:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 12:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 12, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(4 * time.Hour) // 11:00:00
		expected_end := EndOfHour(expected_start, period)                       // 12:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("EndOfHour method for 2 hours period: 14:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 14, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(6 * time.Hour) // 13:00:00
		expected_end := EndOfHour(expected_start, period)                       // 14:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 16:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 16, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(8 * time.Hour) // 15:00:00
		expected_end := EndOfHour(expected_start, period)                       // 16:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 18:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 18, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(10 * time.Hour) // 17:00:00
		expected_end := EndOfHour(expected_start, period)                        // 18:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 20:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 20, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(12 * time.Hour) // 19:00:00
		expected_end := EndOfHour(expected_start, period)                        // 20:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 22:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 22, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(14 * time.Hour) // 21:00:00
		expected_end := EndOfHour(expected_start, period)                        // 22:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 00:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 0, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day-1, 23, 0, 0, 0, thai_time.Location) // 23:00:00
		expected_end := EndOfHour(expected_start, period)                                // 00:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 02:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 2, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 1, 0, 0, 0, thai_time.Location) // 01:00:00
		expected_end := EndOfHour(expected_start, period)                             // 02:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 04:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 4, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 3, 0, 0, 0, thai_time.Location) // 03:00:00
		expected_end := EndOfHour(expected_start, period)                             // 04:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 2 hours period: 06:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 6, 12, 0, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 5, 0, 0, 0, thai_time.Location) // 05:00:00
		expected_end := EndOfHour(expected_start, period)                             // 06:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTime_GetUnixBnStartAndEndOfPeriodHours_4HoursPeriod(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	period := 4
	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 4 hours period: 08:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 8, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time) // 07:00:00
		expected_end := EndOfHour(expected_start, period)    // 10:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 4 hours period: 12:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 12, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(4 * time.Hour) // 11:00:00
		expected_end := EndOfHour(expected_start, period)                       // 14:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 4 hours period: 16:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 16, 12, 56, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(8 * time.Hour) // 15:00:00
		expected_end := EndOfHour(expected_start, period)                       // 18:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 4 hours period: 20:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 20, 12, 56, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(12 * time.Hour) // 19:00:00
		expected_end := EndOfHour(expected_start, period)                        // 22:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 4 hours period: 00:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 0, 12, 56, 0, thai_time.Location)
		expected_start := time.Date(year, month, day-1, 23, 0, 0, 0, thai_time.Location) // 23:00:00
		expected_end := EndOfHour(expected_start, period)                                // 02:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 4 hours period: 04:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 4, 12, 56, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 3, 0, 0, 0, thai_time.Location) // 03:00:00
		expected_end := EndOfHour(expected_start, period)                             // 06:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTime_GetUnixBnStartAndEndOfPeriodHours_6HoursPeriod(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	period := 6
	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 6 hours period: 08:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 8, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time) // 07:00:00
		expected_end := EndOfHour(expected_start, period)    // 12:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 6 hours period: 13:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 13, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(6 * time.Hour) // 13:00:00
		expected_end := EndOfHour(expected_start, period)                       // 18:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 6 hours period: 19:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 19, 12, 56, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(12 * time.Hour) // 19:00:00
		expected_end := EndOfHour(expected_start, period)                        // 00:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 6 hours period: 00:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 0, 12, 56, 0, thai_time.Location)
		// expected_start := time.Date(2024, time.Month(12), 12, 19, 0, 0, 0, thai_time.Location) // 19:00:00
		expected_start := time.Date(2024, time.Month(12), 12, 19, 0, 0, 0, time.UTC).Add(-7 * time.Hour) // 19:00:00
		expected_end := EndOfHour(expected_start, period)                                                // 00:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 6 hours period: 03:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 3, 12, 56, 0, thai_time.Location)
		expected_start := time.Date(year, month, day, 1, 0, 0, 0, thai_time.Location) // 01:00:00
		expected_end := EndOfHour(expected_start, 6)                                  // 06:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, 6)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTime_GetUnixBnStartAndEndOfPeriodHours_8HoursPeriod(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	period := 8
	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 8 hours period: 08:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 8, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time) // 07:00:00
		expected_end := EndOfHour(expected_start, period)    // 14:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 8 hours period: 16:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 16, 12, 56, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(8 * time.Hour) // 15:00:00
		expected_end := EndOfHour(expected_start, period)                       // 22:59:59

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 8 hours period: 00:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 0, 12, 56, 0, thai_time.Location)
		expected_start := time.Date(year, month, day-1, 23, 0, 0, 0, thai_time.Location) // this work as well
		expected_end := EndOfHour(expected_start, period)

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTime_GetUnixBnStartAndEndOfPeriodHours_12HoursPeriod(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	period := 12
	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 12 hours period: 08:12 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 8, 12, 0, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time)
		expected_end := EndOfHour(expected_start, period)

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetUnixBnStartAndEndOfPeriodHours method for 12 hours period: 20:12:56 asia/bangkok", func(t *testing.T) {
		thai_time, _ := thaitime.NewThaiTime()
		_t := time.Date(year, month, day, 20, 12, 56, 0, thai_time.Location)
		expected_start := bnNewDayForThaiTime(_t, thai_time).Add(12 * time.Hour)
		expected_end := EndOfHour(expected_start, period)

		start, end, err := GetUnixBnStartAndEndOfPeriodHours(_t, thai_time, period)

		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}
