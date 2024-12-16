package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetInterval(t *testing.T) {
	t.Run("1h", func(t *testing.T) {
		period, unit, err := GetInterval("1h")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		hour := time.Duration(1) * time.Hour
		assert.Equal(t, period, 1)
		assert.Equal(t, unit, hour)
		assert.Equal(t, unit, Hour)
		assert.Nil(t, err)
	})

	t.Run("12h", func(t *testing.T) {
		period, unit, err := GetInterval("12h")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		hour := time.Duration(1) * time.Hour
		assert.Equal(t, period, 12)
		assert.Equal(t, unit, hour)
		assert.Equal(t, unit, Hour)
		assert.Nil(t, err)
	})

	t.Run("1d", func(t *testing.T) {
		period, unit, err := GetInterval("1d")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		day := time.Duration(1) * time.Hour * 24
		assert.Equal(t, period, 1)
		assert.Equal(t, unit, day)
		assert.Equal(t, unit, Day)
		assert.Nil(t, err)
	})

	t.Run("1w", func(t *testing.T) {
		period, unit, err := GetInterval("1w")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		week := time.Duration(1) * time.Hour * 24 * 7
		assert.Equal(t, period, 1)
		assert.Equal(t, unit, week)
		assert.Equal(t, unit, Week)
		assert.Nil(t, err)
	})
}
