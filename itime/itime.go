package itime

import (
	"strconv"
	"time"
)

const (
	FormatDateOnlyNumber = "20060102"
	FormatTimeOnlyNumber = "150405"
)

func Now() time.Time {
	return time.Now()
}

func AddDays(t time.Time, days int) time.Time {
	return t.Add(time.Hour * 24 * time.Duration(days))
}

func DateOnly(t time.Time) time.Time {
	result, _ := time.Parse(time.DateOnly, t.Format(time.DateOnly))
	return result
}

func DateOnlyNumber(t time.Time) int {
	format := t.Format(FormatDateOnlyNumber)
	atoi, _ := strconv.Atoi(format)
	return atoi
}

func TimeOnlyNumber(t time.Time) int {
	format := t.Format(FormatTimeOnlyNumber)
	atoi, _ := strconv.Atoi(format)
	return atoi
}

func Tomorrow() time.Time {
	return AddDays(Now(), 1)
}
func Yesterday() time.Time {
	return AddDays(Now(), -1)
}

func TodayDateOnly() time.Time {
	return DateOnly(Now())
}

func TomorrowDateOnly() time.Time {
	return DateOnly(Tomorrow())
}

func YesterdayDateOnly() time.Time {
	return DateOnly(Yesterday())
}

func TodayDateOnlyNumber() int {
	return DateOnlyNumber(Now())
}

func TomorrowDateOnlyNumber() int {
	return DateOnlyNumber(Tomorrow())
}

func YesterdayDateOnlyNumber() int {
	return DateOnlyNumber(Yesterday())
}

func NowTimeOnlyNumber() int {
	return TimeOnlyNumber(Now())
}
