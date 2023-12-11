package ijson

import (
	"fmt"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestParseObject(t *testing.T) {
	user, err := BytesToObj[*User]([]byte(`{"name":"name"}`))
	fmt.Println(user, err)
}

func TestStringToObj(t *testing.T) {
	obj, err := StringToObj[User](`{"name":"name"}`)
	fmt.Println("obj", ObjToString(obj), err)

	obj1, err := StringToObj[*User](`{"name":"name"}`)
	fmt.Println("obj1", ObjToString(obj1), err)

	obj2, err := StringToObj[string](`{"name":"name"}`)
	fmt.Println("obj2", ObjToString(obj2), err)

	obj3, err := StringToObj[map[string]any](`{"name":"name"}`)
	fmt.Println("obj3", ObjToString(obj3), err)

	obj4, err := Convert[map[string]any](User{
		Name: "33333",
		Age:  10,
	})
	fmt.Println("obj4", ObjToString(obj4), err)

	u := map[string]any{"name": "1111", "age": 1}
	obj5, err := Convert[*User](u)
	fmt.Println("obj5", ObjToString(obj5), err)

}
