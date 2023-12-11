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
