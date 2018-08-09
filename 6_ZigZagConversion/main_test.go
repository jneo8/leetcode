package zigzagconversion

import "testing"

func Test_main(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{

        {
            name: "case1",
            args: args{
                s: "PAYPALISHIRING",
                numRows: 4,
            },
            want: "PINALSIGYAHRPI",
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("main() = %v, want %v", got, tt.want)
			}
		})
	}
}
