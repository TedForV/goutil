package set

// Union union int32 sets
func Union(data ...[]int32) []int32 {
	m := make(map[int32]bool)
	for _, item := range data {
		for _, v := range item {
			m[v] = true
		}
	}
	result := make([]int32, len(m))
	index := 0
	for i := range m {
		result[index] = i
		index++
	}
	return result
}
