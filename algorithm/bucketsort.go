package algorithm

import (
	"sort"
)

// BucketSort is bucket sort algorithm
func BucketSort(data []float64) []float64 {
	if len(data) == 0 {
		return data
	}

	max, min := data[0], data[0]

	for _, v := range data {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	d := max - min

	bucketNum := len(data)

	var buckets [][]float64
	for i := 0; i < bucketNum; i++ {
		buckets = append(buckets, []float64{})
	}

	for _, v := range data {
		pos := int(((v - min) / d) * float64(bucketNum-1))
		buckets[pos] = append(buckets[pos], v)
	}

	result := make([]float64, 0, len(data))

	for _, v := range buckets {
		if len(v) == 0 {
			continue
		}
		if len(v) == 1 {
			result = append(result, v[0])
			continue
		}
		sort.Float64s(v)

		result = append(result, v...)

	}
	return result
}
