package utils

import "time"

func GetBinanceTimestamp() int64 {
	return time.Now().Unix() * 1000
}
