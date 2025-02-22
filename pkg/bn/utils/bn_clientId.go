package utils

import (
	"fmt"
	"time"
)

// default client id with running number from symbol
func BinanceDefaultClientID(symbol string, position_side string, counting int64) string {
	return fmt.Sprintf("%s_%s_%d", symbol, position_side, counting)
}

// default client id without running number from symbol but got unix nano instead
func BinanceDefaultClientIDV2(symbol string, position_side string) string {
	return fmt.Sprintf("%s_%s_%d", symbol, position_side, time.Now().UnixNano())
}
