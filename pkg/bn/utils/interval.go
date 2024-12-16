package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	Minute = time.Duration(1) * time.Minute
	Hour   = time.Duration(1) * time.Hour
	Day    = time.Duration(1) * time.Hour * 24
	Week   = time.Duration(1) * time.Hour * 24 * 7
	// Month  = time.Duration(1) * time.Hour * 24 * 30
)

func GetInterval(interval string) (int, time.Duration, error) {
	interval = strings.ToLower(interval)
	period := interval[:len(interval)-1]
	_period, err := strconv.Atoi(period)
	if err != nil {
		return 0, 0, err
	}
	_unit := interval[len(interval)-1]
	switch _unit {
	case 'm':
		return _period, Minute, nil
	case 'h':
		return _period, Hour, nil
	case 'd':
		return _period, Day, nil
	case 'w':
		return _period, Week, nil
	// case 'M':
	// 	return _period, Month, nil
	default:
		return 0, 0, errors.New("invalid interval")
	}
}
