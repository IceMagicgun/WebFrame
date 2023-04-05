package time_tool

import "time"

func Now() int64 {
	return time.Now().Unix()
}

func NowText() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
