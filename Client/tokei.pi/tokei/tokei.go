package tokei

import "time"

func GetCurrentWeekdayString() string {
	currentTime := time.Now()
	return currentTime.Format("Monday")
}

func GetCurrentDateString() string {
	currentTime := time.Now()
	return currentTime.Format("01-02-2006")
}

func GetCurrentTimeString() string {
	currentTime := time.Now()
	return currentTime.Format("15:04:05")
}

func GetTimeStringFromTimeStamp(timestamp int64) string {
	currentTime := time.Unix(int64(timestamp), 0)
	return currentTime.Format("15:04:05")
}

func GetTimeStringFromTimeStampLowPrecision(timestamp int64) string {
	currentTime := time.Unix(int64(timestamp), 0)
	return currentTime.Format("15:04")
}
