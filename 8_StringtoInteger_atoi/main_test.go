package string_to_integer_atoi

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
        {
            name: "case1",
            args: args{str: "42"},
            want: 42,
        },
        {
            name: "case2",
            args: args{str: "   -42"},
            want: -42,
        },

        {
            name: "case3",
            args: args{str: "words and 987"},
            want: 0,
        },
        {
            name: "case4",
            args: args{str: "3.14159"},
            want: 3,
        },
        {
            name: "case5",
            args: args{str: "+-2"},
            want: 0,
        },
        {
            name: "case6",
            args: args{str: "    +0  123"},
            want: 0,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
