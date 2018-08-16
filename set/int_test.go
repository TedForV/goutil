package set

import (
	"testing"
)

func TestUnion(t *testing.T) {
	type args struct {
		data [][]int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "union 1",
			args: args{
				data: [][]int32{[]int32{1, 2, 3, 4, 5}, []int32{2, 3, 4, 5, 6, 7}, []int32{1, 5, 9, 0, 10}},
			},
			want: []int32{0, 1, 2, 3, 4, 5, 6, 7, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.data...); len(got) != len(tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}
