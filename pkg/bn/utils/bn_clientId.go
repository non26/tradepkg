package utils

import "fmt"

func BinanceDefaultClientID(symbol string, position_side string, counting int) string {
	return fmt.Sprintf("%s_%s_%d", symbol, position_side, counting)
}
