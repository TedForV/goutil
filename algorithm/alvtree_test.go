package algorithm

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	type args struct {
		data []int32
		root *ALVNode
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				data: []int32{18, 14, 20, 12, 16, 1, -1},
				root: nil,
			},
			want: []int32{-1, 1, 12, 14, 16, 18, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.data {
				tt.args.root = Insert(v, tt.args.root)
			}
			got := PreorderTraversal(tt.args.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
			// if got := Insert(tt.args.data, tt.args.root); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Insert() = %v, want %v", got, tt.want)
			// }
		})
	}
}
