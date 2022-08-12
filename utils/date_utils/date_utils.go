package date_utils

import "time"

const (
	apiDateLayout = "2006-02-01T15:05:05Z"
)

// GetNow Get the current date.
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString Get the current date as string.
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
