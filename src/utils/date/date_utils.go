package dateutils

import "time"

const apiDateLayout = "2006-01-02T15:01"

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetTodayDateInString() string {
	return GetNow().Format(apiDateLayout)
}
