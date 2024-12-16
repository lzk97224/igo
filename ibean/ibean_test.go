package ibean

import (
	"fmt"
	"testing"
)

func TestCopyBean(t *testing.T) {
	type School struct {
		Name    string
		Address string
	}
	type Home struct {
		Name    string
		Address string
	}
	type People struct {
		Name    string
		Age     int
		School  *School
		School2 *School
		Home    Home
	}

	p1 := People{
		Name:    "san",
		Age:     10,
		School2: nil,
		School: &School{
			Name:    "sss",
			Address: "sss",
		},
		Home: Home{
			Name:    "xxx",
			Address: "xxx",
		},
	}
	p2 := new(People)
	err := CopyBean(&p1, p2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(p2)

}
