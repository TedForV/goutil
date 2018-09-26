package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name   string
		args   args
		result []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				arr: []int{2, 4, 6, 8, 1, 3, 5, 7, 9, 0},
			},
			result: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "2",
			args: args{
				arr: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			result: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			result: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "4",
			args: args{
				arr: []int{1},
			},
			result: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			assert.Equal(t, tt.result, tt.args.arr)
		})
	}
}
