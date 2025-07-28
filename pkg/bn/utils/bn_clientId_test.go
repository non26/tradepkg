package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinanceDefaultClientID_success(t *testing.T) {
	symbol := "BTCUSDT"
	position_side := "LONG"
	counting := int64(123)
	client_id := BinanceDefaultClientID(symbol, position_side, counting)

	expected_client_id := "BTCUSDT_LONG_123"
	assert.Equal(t, expected_client_id, client_id)
}

func TestBinanceDefaultClientIDV2_success(t *testing.T) {
	symbol := "BTCUSDT"
	position_side := "LONG"
	client_id := BinanceDefaultClientIDV2(symbol, position_side)

	assert.True(t, strings.Contains(client_id, symbol+"_"+position_side))
}
