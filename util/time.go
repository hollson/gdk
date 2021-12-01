// +build linux darwin arm

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

// Zone 获取时区标识与时区值
//  常用时区名称：
//    Asia/Shanghai:       上海
//    Asia/Taipei:         台北
//    Asia/Seoul:          首尔
//    Asia/Tokyo:          东京
//    Asia/Singapore:      新加坡
//    Asia/Jerusalem:      耶路撒冷
//    Asia/Kolkata:        加尔各答(印度)
//    America/Los_Angeles: 洛杉矶
//    America/New_York:    纽约
//    Europe/Dublin:       都柏林
//    Europe/London:       伦敦
//  更多参考：(🔥源码才是最好的参考文档)
//    cat /private/etc/localtime
//    cat etc/localtime
//    ls usr/share/zoneinfo
//    https://www.zeitverschiebung.net/cn/city/2643743
func Zone(loc *time.Location) (name string, zone int) {
	name, zone = ShiftZone(loc).Zone()
	return name, zone / 3600
}

// DateShrink 只包含年月日的时间(即：抹掉时分秒)
func DateShrink(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local)
}

// ToMonday 转到以tar为坐标的周一时间点,tar默认为当前时间
func ToMonday(tar ...time.Time) time.Time {
	t := time.Now()
	if len(tar) > 0 {
		t = tar[0]
	}
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	return t.AddDate(0, 0, offset)
}

// ShiftZone 时区转换/换置(即:将from时间转换为toLoc时区的时间)
func ShiftZone(toLoc *time.Location, from ...time.Time) time.Time {
	if len(from) > 0 {
		return from[0].In(toLoc)
	}
	return time.Now().In(toLoc)
}
