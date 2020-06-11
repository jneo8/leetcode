package main

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				nums:   []int{0, 1, 2},
				target: 3,
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				nums:   []int{0, 2, 1, -3},
				target: 1,
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				nums:   []int{89, -13, 38, -39, -58, 32, 9, 97, -31, 13, 99, -100, -7, 32, 13, -32, -12, 59, 43, -62, 89, -8, 15, 41, 31, -48, -50, -80, -82, -22, -60, 29, 32, 88, 17, -76, 10, 4, 55, 4, 65, 83, -100, 5, 6, 75, -13, 4, 57, -53, 8, 80, -79, -76, 61, -52, 100, 74, -96, 79, -72, -68, -56, -29, 98, -52, -77, 90, -48, 21, 78, -1, -6, 8, 19, 69, -39, -74, -42, -39, -86, 10, 55, -84, -97, -79, 76, 84, -87, -49, -44, 28, -88, -9, 71, 34, -72, 85, 9, -96, -4, -79, -52, 95, 20, -76, 94, 80, -38, -76, -70, -98, -19, 34, -82, 67, 36, -98, 74, -93, 1, -92, -4, 42, -26, -16, 27, 3, -61, 41, 66, -60, 85, -21, -15, 52, 14, 78, -80, 60, -36, -35, -56, -76, -8, -25, 56, -7, -8, -31, 8, 40, -25, -40, 2, -44, 7},
				target: -71,
			},
			want: -71,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
