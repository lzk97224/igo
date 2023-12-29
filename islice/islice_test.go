package islice

import (
	"fmt"
	"testing"
)

type People struct {
	Age  int
	Name string
}

func TestSortByKeySlice(t *testing.T) {

	slice := SortByKeySlice([]People{{
		Age:  1,
		Name: "1",
	}, {
		Age:  3,
		Name: "3",
	}, {
		Age:  2,
		Name: "2",
	}}, []int{2, 1, 3}, func(people People) int {
		return people.Age
	})
	fmt.Println(slice)
}

func TestToMap(t *testing.T) {
	toMap := ToMap[int, int, int]([]int{1, 2, 3}, func(i int) int {
		return i
	}, func(i int) int {
		return i
	})

	ToMapByKey([]int{1, 2, 3}, func(t int) int {
		return t
	})
	fmt.Println(toMap)
}
