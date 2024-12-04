package utils

import (
	"strings"

	positionconstant "github.com/non26/tradepkg/pkg/bn/position_constant"
)

func IsBuyCrypto(side string, position_side string) bool {
	side = strings.ToUpper(side)
	position_side = strings.ToUpper(position_side)
	switch {
	case side == positionconstant.BUY && position_side == positionconstant.LONG:
		return true
	case side == positionconstant.SELL && position_side == positionconstant.SHORT:
		return true
	default:
		return false
	}
}

func IsSellCrypto(side string, position_side string) bool {
	side = strings.ToUpper(side)
	position_side = strings.ToUpper(position_side)
	switch {
	case side == positionconstant.SELL && position_side == positionconstant.LONG:
		return true
	case side == positionconstant.BUY && position_side == positionconstant.SHORT:
		return true
	default:
		return false
	}
}
