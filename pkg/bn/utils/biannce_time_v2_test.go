package utils

import (
	"testing"
	"time"

	"github.com/non26/tradepkg/pkg/bn/thaitime"
	"github.com/stretchr/testify/assert"
)

func TestBinanceTimeV2_GetBnTimeStartHourAndEndHour_2hoursPeriod_success(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	minute := 12
	second := 57
	period := 2
	thai_time, _ := thaitime.NewThaiTime()

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_00:00:00_01:59:59_UTC_at_08:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 8, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 1, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_02:00:00_03:59:59_UTC_at_10:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 10, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 2, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 3, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_04:00:00_05:59:59_UTC_at_12:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 12, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 4, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 5, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_06:00:00_07:59:59_UTC_at_14:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 14, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 6, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 7, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_08:00:00_09:59:59_UTC_at_16:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 16, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 8, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 9, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_10:00:00_11:59:59_UTC_at_18:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 18, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 10, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 11, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_12:00:00_13:59:59_UTC_at_20:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 20, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 13, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_14:00:00_15:59:59_UTC_at_22:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 22, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 14, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 15, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_16:00:00_17:59:59_UTC_at_00:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day+1, 0, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 16, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 17, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_2hoursPeriod_18:00:00_19:59:59_UTC_at_02:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day+1, 2, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 18, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 19, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTimeV2_GetBnTimeStartHourAndEndHour_3hoursPeriod_success(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	minute := 12
	second := 57
	period := 3
	thai_time, _ := thaitime.NewThaiTime()

	t.Run("GetBnTimeStartHourAndEndHour_3hoursPeriod_00:00:00_02:59:59_UTC_at_08:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 8, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 2, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_3hoursPeriod_03:00:00_05:59:59_UTC_at_11:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 10, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 3, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 5, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_3hoursPeriod_06:00:00_08:59:59_UTC_at_14:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 14, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 6, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 8, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_3hoursPeriod_09:00:00_11:59:59_UTC_at_17:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 17, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 9, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 11, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_3hoursPeriod_12:00:00_14:59:59_UTC_at_20:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 20, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 14, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_3hoursPeriod_15:00:00_17:59:59_UTC_at_23:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 23, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 15, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 17, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_3hoursPeriod_18:00:00_20:59:59_UTC_at_02:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day+1, 2, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 18, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 20, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_3hoursPeriod_21:00:00_23:59:59_UTC_at_05:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day+1, 5, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 21, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTimeV2_GetBnTimeStartHourAndEndHour_4hoursPeriod_success(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	minute := 12
	second := 57
	period := 4
	thai_time, _ := thaitime.NewThaiTime()
	t.Run("GetBnTimeStartHourAndEndHour_4hoursPeriod_00:00:00_03:59:59_UTC_at_08:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 8, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 3, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_4hoursPeriod_04:00:00_07:59:59_UTC_at_12:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 12, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 4, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 7, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_4hoursPeriod_08:00:00_11:59:59_UTC_at_16:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 16, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 8, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 11, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_4hoursPeriod_12:00:00_15:59:59_UTC_at_20:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 20, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 15, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_4hoursPeriod_16:00:00_19:59:59_UTC_at_00:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day+1, 0, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 16, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 19, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_4hoursPeriod_20:00:00_23:59:59_UTC_at_04:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day+1, 4, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 20, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTimeV2_GetBnTimeStartHourAndEndHour_6hoursPeriod_success(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	minute := 12
	second := 57
	period := 6
	thai_time, _ := thaitime.NewThaiTime()

	t.Run("GetBnTimeStartHourAndEndHour_6hoursPeriod_00:00:00_05:59:59_UTC_at_08:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 8, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 5, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_6hoursPeriod_06:00:00_11:59:59_UTC_at_14:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 14, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 6, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 11, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_6hoursPeriod_12:00:00_17:59:59_UTC_at_20:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 20, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 17, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTimeV2_GetBnTimeStartHourAndEndHour_8hoursPeriod_success(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	minute := 12
	second := 57
	period := 8
	thai_time, _ := thaitime.NewThaiTime()

	t.Run("GetBnTimeStartHourAndEndHour_8hoursPeriod_00:00:00_07:59:59_UTC_at_08:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 8, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 7, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_8hoursPeriod_08:00:00_15:59:59_UTC_at_16:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 16, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 8, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 15, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_8hoursPeriod_16:00:00_23:59:59_UTC_at_24:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day+1, 24, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 16, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}

func TestBinanceTimeV2_GetBnTimeStartHourAndEndHour_12hoursPeriod_success(t *testing.T) {
	year := 2024
	month := time.Month(12)
	day := 13
	minute := 12
	second := 57
	period := 12
	thai_time, _ := thaitime.NewThaiTime()

	t.Run("GetBnTimeStartHourAndEndHour_12hoursPeriod_00:00:00_11:59:59_UTC_at_08:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 8, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 11, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})

	t.Run("GetBnTimeStartHourAndEndHour_12hoursPeriod_12:00:00_23:59:59_UTC_at_20:12:57_asia/bangkok", func(t *testing.T) {
		_thai_time := time.Date(year, month, day, 20, minute, second, 0, thai_time.Location)
		bn_time := NewBinanceTime(_thai_time)

		expected_start := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
		expected_end := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)

		start, end, err := bn_time.GetBnTimeStartHourAndEndHour(period)
		assert.Nil(t, err)
		assert.Equal(t, start.Unix(), expected_start.Unix())
		assert.Equal(t, end.Unix(), expected_end.Unix())
	})
}
