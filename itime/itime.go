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

func NowUnix() int64 {
	return Now().Unix()
}
func NowUnixMilli() int64 {
	return Now().UnixMilli()
}

func AddYear(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}
func AddMonth(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}
func AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Hour * time.Duration(hours))
}
func AddMinute(t time.Time, minutes int) time.Time {
	return t.Add(time.Minute * time.Duration(minutes))
}
func AddSecond(t time.Time, seconds int) time.Time {
	return t.Add(time.Second * time.Duration(seconds))
}

func DateOnly(t time.Time) time.Time {
	result, _ := time.ParseInLocation(time.DateOnly, t.Format(time.DateOnly), time.Local)
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
