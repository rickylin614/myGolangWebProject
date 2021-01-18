package utils

import "time"

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func TimeToString(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(TimeFormat)
}
