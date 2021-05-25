package goox

import "time"

// 自定义Json成员类型
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
