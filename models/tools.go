package models

import (
	"time"
)

func UnixToTime(timestamp int64) string {
	// Convert Unix timestamp to human-readable format
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
