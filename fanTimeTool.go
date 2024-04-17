package fanTool

import (
	"time"
)

// ThisMorningUnixTime 当天0点的时间戳
func ThisMorningUnixTime() int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

// YesterdayMorningUnixTime 昨天0点的时间戳
func YesterdayMorningUnixTime() int64 {
	t := time.Now().AddDate(0, 0, -1)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

// StringToTime 时间戳转时间字符串
func StringToTime(s string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	if err == nil {
		return t
	}
	t, err = time.ParseInLocation("2006-01-02", s, loc)
	if err == nil {
		return t
	}
	t, err = time.ParseInLocation("2006/01/02 15:04:05", s, loc)
	if err == nil {
		return t
	}
	t, err = time.ParseInLocation("2006/01/02", s, loc)
	if err == nil {
		return t
	}
	t, err = time.ParseInLocation("20060102", s, loc)
	if err == nil {
		return t
	}
	return t
}

// StringToUnixTime 时间字符串转时间戳
func StringToUnixTime(s string) int64 {
	return StringToTime(s).Unix()
}
