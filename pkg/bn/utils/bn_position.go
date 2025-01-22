package utils

import (
	"strings"

	constant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

func IsBuyPosition(side string, position_side string) bool {
	side = strings.ToUpper(side)
	position_side = strings.ToUpper(position_side)
	switch {
	case side == constant.BUY && position_side == constant.LONG:
		return true
	case side == constant.SELL && position_side == constant.SHORT:
		return true
	default:
		return false
	}
}

func IsSellPosition(side string, position_side string) bool {
	side = strings.ToUpper(side)
	position_side = strings.ToUpper(position_side)
	switch {
	case side == constant.SELL && position_side == constant.LONG:
		return true
	case side == constant.BUY && position_side == constant.SHORT:
		return true
	default:
		return false
	}
}

func IsLongPosition(position_side string) bool {
	return position_side == constant.LONG
}

func IsShortPosition(position_side string) bool {
	return position_side == constant.SHORT
}
