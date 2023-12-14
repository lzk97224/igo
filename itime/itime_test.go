package itime

import (
	"fmt"
	"testing"
)

func TestTodayDate(t *testing.T) {
	fmt.Println(TodayDateOnly())
}

func TestTomorrowDate(t *testing.T) {
	fmt.Println(TomorrowDateOnly())
}

func TestYesterdayDate(t *testing.T) {
	fmt.Println(YesterdayDateOnly())
}

func TestNowTimeOnlyNumber(t *testing.T) {
	fmt.Println(NowTimeOnlyNumber())
}

func TestAddYear(t *testing.T) {
	fmt.Println(AddYear(Now(), 2))
	fmt.Println(AddYear(Now(), 3))
}

func TestAddMonth(t *testing.T) {
	fmt.Println(AddMonth(Now(), 1))
	fmt.Println(AddMonth(Now(), 2))
	fmt.Println(AddMonth(Now(), 3))
	fmt.Println(AddMonth(Now(), 4))
}

func TestAddDayss(t *testing.T) {
	now := Now()
	fmt.Println(AddDays(now, 0))
	fmt.Println(AddDays(now, 3))
	fmt.Println(AddDays(now, 4))
	fmt.Println(AddDays(now, 5))
	fmt.Println(AddDays(now, 6))
	fmt.Println(AddDays(now, 20))

	fmt.Println(AddDays(now, 0))
	fmt.Println(AddDays(now, 3))
	fmt.Println(AddDays(now, 4))
	fmt.Println(AddDays(now, 5))
	fmt.Println(AddDays(now, 6))
	fmt.Println(AddDays(now, 20))
}
