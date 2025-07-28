package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBuyPosition(t *testing.T) {

	t.Run("For Long Position", func(t *testing.T) {
		side := "BUY"
		position_side := "LONG"
		assert.True(t, IsBuyPosition(side, position_side))
	})

	t.Run("For Short Position", func(t *testing.T) {
		side := "SELL"
		position_side := "SHORT"
		assert.True(t, IsBuyPosition(side, position_side))
	})

}

func TestIsSellPosition(t *testing.T) {
	t.Run("For Long Position", func(t *testing.T) {
		side := "SELL"
		position_side := "LONG"
		assert.True(t, IsSellPosition(side, position_side))
	})

	t.Run("For Short Position", func(t *testing.T) {
		side := "BUY"
		position_side := "SHORT"
		assert.True(t, IsSellPosition(side, position_side))
	})
}

func TestIsLongPosition(t *testing.T) {
	t.Run("For Long Position", func(t *testing.T) {
		position_side := "LONG"
		assert.True(t, IsLongPosition(position_side))
	})

	t.Run("For Short Position", func(t *testing.T) {
		position_side := "SHORT"
		assert.False(t, IsLongPosition(position_side))
	})
}

func TestIsShortPosition(t *testing.T) {
	t.Run("For Short Position", func(t *testing.T) {
		position_side := "SHORT"
		assert.True(t, IsShortPosition(position_side))
	})

	t.Run("For Long Position", func(t *testing.T) {
		position_side := "LONG"
		assert.False(t, IsShortPosition(position_side))
	})
}

func TestToSellSideBy(t *testing.T) {
	t.Run("sell Long Position", func(t *testing.T) {
		position_side := "LONG"
		assert.Equal(t, ToSellSideBy(position_side), "SELL")
	})

	t.Run("sell Short Position", func(t *testing.T) {
		position_side := "SHORT"
		assert.Equal(t, ToSellSideBy(position_side), "BUY")
	})
}

func TestToBuySideBy(t *testing.T) {
	t.Run("buy Long Position", func(t *testing.T) {
		position_side := "LONG"
		assert.Equal(t, ToBuySideBy(position_side), "BUY")
	})

	t.Run("buy Short Position", func(t *testing.T) {
		position_side := "SHORT"
		assert.Equal(t, ToBuySideBy(position_side), "SELL")
	})
}
