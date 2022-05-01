package gdk

import (
	"time"
)

const (
	DefaultDateFormat = "2006-01-02"
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

// JsonTime 自定义Json成员类型
type JsonTime time.Time

func (p JsonTime) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0, 24)
	data = append(data, '"')
	data = time.Time(p).AppendFormat(data, "2006-01-02 15:04:05")
	data = append(data, '"')
	return data, nil
}

func (p JsonTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02 15:04:05"`, string(b))
	p = JsonTime(t)
	return err
}

func (p JsonTime) String() string {
	return time.Time(p).Format("2006-01-02 15:04:05")
}

// Deprecated: 推荐使用 time.Unix(0, 0)
//  1970年1月1日(08:00:00GMT)
func Time1970() time.Time {
	return time.Unix(0, 0)
}

// ParseTime 时间转换 (推荐使用 time.ParseInLocation)
//  将"2006-01-02 15:04:05"格式的字符串转换为时间. 参数错误时返回 1970-01-01(08:00:00GMT).
func ParseTime(s string) time.Time {
	local, _ := time.LoadLocation("Asia/Shanghai")
	ret, err := time.ParseInLocation(DefaultTimeFormat, s, local)
	if err != nil {
		return time.Unix(0, 0)
	}
	return ret
}

// ParseDate 日前转换 (推荐使用 time.ParseInLocation)
//  将"2006-01-02"格式的字符串转换为时间. 参数错误时返回 1970-01-01(08:00:00GMT).
func ParseDate(s string) time.Time {
	local, _ := time.LoadLocation("Asia/Shanghai")
	ret, err := time.ParseInLocation(DefaultDateFormat, s, local)
	if err != nil {
		return time.Unix(0, 0)
	}
	return ret
}
