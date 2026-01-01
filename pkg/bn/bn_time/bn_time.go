package bntime

import (
	"time"
)

func GetBinanceTimestamp() int64 {
	return time.Now().Unix() * 1000
}

func GetSpecificBnTimestamp(t *time.Time) int64 {
	return t.Unix() * 1000
}

func GetDBTime() string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc).Format(time.RFC3339)
}
