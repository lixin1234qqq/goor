package utils

import "time"

func CurrentISO8601Time() string {
	return time.Now().Format("2006-01-02T15:04:05-0700")
}
