package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpAdjust(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				arr: []int{1, 3, 2, 6, 5, 7, 8, 9, 10, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpAdjust(tt.args.arr)
			assert.Equal(t, []int{0, 1, 2, 6, 3, 7, 8, 9, 10, 5}, tt.args.arr)
		})
	}
}

func TestDownAdjust(t *testing.T) {
	type args struct {
		arr         []int
		parentIndex int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				arr:         []int{10, 3, 2, 6, 5, 7, 8, 9},
				parentIndex: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DownAdjust(tt.args.arr, tt.args.parentIndex, len(tt.args.arr))
			assert.Equal(t, []int{2, 3, 7, 6, 5, 10, 8, 9}, tt.args.arr)
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				[]int{5, 3, 4, 2, 1, 6, 9, 7, 8, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.arr)
			assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, tt.args.arr)
		})
	}
}
