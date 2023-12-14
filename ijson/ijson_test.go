package ijson

import (
	"fmt"
	"reflect"
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

func TestStringToObj1(t *testing.T) {
	type args struct {
		jsonStr string
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[User]{
		{name: "", args: args{jsonStr: ""}, want: User{}, wantErr: true},
		{name: "", args: args{jsonStr: "{}"}, want: User{}, wantErr: false},
		{name: "", args: args{jsonStr: `{"age":"23"}`}, want: User{}, wantErr: true},
		{name: "", args: args{jsonStr: "[]"}, want: User{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToObj[User](tt.args.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToObj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToObj() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToObj2(t *testing.T) {
	type args struct {
		jsonStr string
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[*User]{
		{name: "", args: args{jsonStr: ""}, want: nil, wantErr: true},
		{name: "", args: args{jsonStr: "{}"}, want: &User{}, wantErr: false},
		{name: "", args: args{jsonStr: `{"age":"sdfj"}`}, want: &User{}, wantErr: true},
		{name: "", args: args{jsonStr: "[]"}, want: &User{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToObj[*User](tt.args.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToObj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToObj() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToObj3(t *testing.T) {
	type args struct {
		jsonStr string
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[map[string]any]{
		{name: "", args: args{jsonStr: ""}, want: nil, wantErr: true},
		{name: "", args: args{jsonStr: "{}"}, want: map[string]any{}, wantErr: false},
		{name: "", args: args{jsonStr: `{"age":"sdfj"}`}, want: map[string]any{"age": "sdfj"}, wantErr: false},
		{name: "", args: args{jsonStr: "[]"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToObj[map[string]any](tt.args.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToObj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToObj() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToObj4(t *testing.T) {
	type args struct {
		jsonStr string
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[[]*User]{
		{name: "", args: args{jsonStr: ""}, want: nil, wantErr: true},
		{name: "", args: args{jsonStr: "{}"}, want: nil, wantErr: true},
		{name: "", args: args{jsonStr: `{"age":"sdfj"}`}, want: nil, wantErr: true},
		{name: "", args: args{jsonStr: "[]"}, want: []*User{}, wantErr: false},
		{name: "", args: args{jsonStr: `[{"age":"234"}]`}, want: []*User{&User{}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToObj[[]*User](tt.args.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToObj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToObj() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjToBytes(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "", args: args{v: nil}, want: []byte{'n', 'u', 'l', 'l'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ObjToBytes(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObjToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
