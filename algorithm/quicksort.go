package algorithm

// QuickSort for quick sort
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid, i := arr[0], 1
	left, right := 0, len(arr)-1
	for left < right {
		if arr[i] > mid {
			arr[i], arr[right] = arr[right], arr[i]
			right--
		} else {
			arr[i], arr[left] = arr[left], arr[i]
			left++
			i++
		}
	}
	arr[left] = mid
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
