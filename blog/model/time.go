package model

import "time"

const timeFormat = "2006-01-02 15:04:05"
const timeZone = "Asia/Shanghai"

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(b), time.Local)
	if err == nil {
		return err
	}
	*t = Time(now)
	return nil
}
