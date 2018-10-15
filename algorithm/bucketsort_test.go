package algorithm

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				data: []float64{4.12, 6.421, 0.0023, 3.0, 2.123, 8.122, 4.12, 10.09},
			},
			want: []float64{0.0023, 2.123, 3.0, 4.12, 4.12, 6.421, 8.122, 10.09},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketSort(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
