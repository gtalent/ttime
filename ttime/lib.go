package ttime

import (
	"strconv"
	"time"
)

func Tdate() string {
	n := time.Now()
	n = time.Date(n.Year(), n.Month(), n.Day(), 24, 60, 60, 0, time.Local)
	var ny time.Time
	ny = time.Date(n.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	day := (n.Unix() - ny.Unix()) / 86400
	return strconv.FormatInt(day, 10) + ", " + strconv.Itoa(n.Year())
}

func Ttime() string {
	n := time.Now()
	var tsec int64
	{
		t := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local)
		tsec = (n.Unix() - t.Unix()) * 1000
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
	str := strconv.FormatInt(val, 10)
	for len(str) < desiredLen {
		str = "0" + str
	}
	return str
}
