package ibytes

import "testing"

func TestToString(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{[]byte{}}, want: ""},
		{name: "test2", args: args{nil}, want: ""},
		{name: "test3", args: args{[]byte{'1'}}, want: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.src); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
