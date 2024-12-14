package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
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
		return _period, time.Duration(1) * time.Minute, nil
	case 'h':
		return _period, time.Duration(1) * time.Hour, nil
	case 'd':
		return _period, time.Duration(1) * time.Hour * 24, nil
	case 'w':
		return _period, time.Duration(1) * time.Hour * 24 * 7, nil
	default:
		return 0, 0, errors.New("invalid interval")
	}
}
