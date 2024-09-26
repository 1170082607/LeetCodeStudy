package array_and_strings

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,3,4,5,2",
			args: args{
				ratings: []int{1, 3, 4, 5, 2},
			},
			want: 11,
		},
		{
			name: "1,3,2,2,1",
			args: args{
				ratings: []int{1, 3, 2, 2, 1},
			},
			want: 7,
		},
		{
			name: "1,0,2",
			args: args{
				ratings: []int{1, 0, 2},
			},
			want: 5,
		},
		{
			name: "1,2,2",
			args: args{
				ratings: []int{1, 2, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
