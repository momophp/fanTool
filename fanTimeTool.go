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

func AddDaysUnixTime(days int, timeUnix int64) int64 {
	if days == 0 {
		days = 1
	}
	if timeUnix == 0 {
		timeUnix = time.Now().Unix()
	}
	return time.Unix(timeUnix, 0).AddDate(0, 0, days).Unix()
}

func AddMonthsUnixTime(months int, timeUnix int64) int64 {
	if months == 0 {
		months = 1
	}
	if timeUnix == 0 {
		timeUnix = time.Now().Unix()
	}
	return time.Unix(timeUnix, 0).AddDate(0, months, 0).Unix()
}

func AddYearsUnixTime(years int, timeUnix int64) int64 {
	if years == 0 {
		years = 1
	}
	if timeUnix == 0 {
		timeUnix = time.Now().Unix()
	}
	return time.Unix(timeUnix, 0).AddDate(years, 0, 0).Unix()
}

func AddHoursUnixTime(hours int, timeUnix int64) int64 {
	if hours == 0 {
		hours = 1
	}
	if timeUnix == 0 {
		timeUnix = time.Now().Unix()
	}
	return time.Unix(timeUnix, 0).Add(time.Duration(hours) * time.Hour).Unix()
}
