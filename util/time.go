package util

import (
	"time"
)

type TimeTemplate string

const (
	TT1 TimeTemplate = "2006-01-02 15:04:05"
	TT2              = "2006/01/02 15:04:05"
	TT3              = "2006-01-02"
	TT4              = "20060102"
	TT5              = "15:04:05"
)

func StrToTime(timeStr string, template TimeTemplate) (time.Time, error) {
	return time.ParseInLocation(string(template), timeStr, time.Local)
}

// 向前的倒退的周N
func RetroWeekN(n time.Weekday) time.Time {
	now := time.Now()
	offset := int(n - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}


// 获取年月日部分
func TimeOfYMD(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local)
}

// 本周一
func ThisMonday() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

