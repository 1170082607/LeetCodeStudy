package array_and_strings

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
		},
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				k:    3,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    11,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27},
				k:    38,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
		})
	}
}
