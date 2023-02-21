package dates

import "time"

const (
	ApiDateLayout = "2006-01-02T15:04:05Z"
	ApiDBLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(ApiDateLayout)
}

func GetNowDbFormat() string {
	return GetNow().Format(ApiDBLayout)
}
