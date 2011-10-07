package ttime

import (
	"strconv"
	"time"
)

func Tdate() string {
	now := time.LocalTime()
	now.Hour = 24
	now.Minute = 60
	now.Second = 60
	var ny time.Time
	ny.Year = now.Year
	ny.Month = 1
	ny.Day = 1
	ny.Hour = 0
	ny.Minute = 0
	ny.Second = 0
	ny.Zone = "CST"
	day := (now.Seconds() - ny.Seconds()) / 86400
	return strconv.Itoa64(day) + ", " + strconv.Itoa64(now.Year)
}

func Ttime() string {
	now := time.LocalTime()
	var tsec int64
	{
		var t time.Time = *now
		t.Hour = 0
		t.Minute = 0
		t.Second = 0
		tsec = (now.Seconds() - t.Seconds()) * 1000
	}
	ms := int64(86400 * 1000)
	hour := int64((float64(tsec) / float64(ms)) * 10)
	tsec -= int64((ms / 10) * hour)
	minute := int64((float64(tsec) / (float64(ms) / 10)) * 100)
	tsec -= int64((ms / 1000) * minute)
	second := int64((float64(tsec) / (float64(ms) / 100)) * 10000)
	return format(hour, 1) + ":" + format(minute, 2) + ":" + format(second, 3)
}

func format(val int64, desiredLen int) string {
	str := strconv.Itoa64(val)
	for len(str) < desiredLen {
		str = "0" + str
	}
	return str
}
