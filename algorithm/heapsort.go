package algorithm

// UpAdjust upajust the last node (inserted new node situation)
func UpAdjust(arr []int) {
	if len(arr) == 0 || len(arr) == 1 {
		return
	}
	childIndex := len(arr) - 1
	parentIndx := getParentIndex(childIndex)
	for {
		if childIndex == 0 && arr[childIndex] >= arr[parentIndx] {
			break
		}
		arr[childIndex], arr[parentIndx] = arr[parentIndx], arr[childIndex]
		childIndex = parentIndx
		parentIndx = getParentIndex(childIndex)
	}
}

// DownAdjust downadjust the node (replace the root node with last node)
// func DownAdjust(arr []int, parentIndex int) {
// 	length := len(arr)
// 	if length <= parentIndex+1 {
// 		return
// 	}
// 	childIndex := parentIndex*2 + 1
// 	if length < childIndex+1 {
// 		return
// 	}
// 	for {
// 		if length < childIndex+1 {
// 			break
// 		}
// 		if childIndex+2 <= length && arr[childIndex] > arr[childIndex+1] {
// 			childIndex++
// 		}
// 		if arr[parentIndex] <= arr[childIndex] {
// 			break
// 		}
// 		arr[childIndex], arr[parentIndex] = arr[parentIndex], arr[childIndex]
// 		parentIndex = childIndex
// 		childIndex = parentIndex*2 + 1
// 	}
// }

// DownAdjust downadjust the node (replace the root node with last node)
func DownAdjust(arr []int, parentIndex int, validLength int) {
	if validLength <= parentIndex+1 {
		return
	}
	childIndex := parentIndex*2 + 1
	if validLength < childIndex+1 {
		return
	}
	for {
		if validLength < childIndex+1 {
			break
		}
		if childIndex+2 <= validLength && arr[childIndex] > arr[childIndex+1] {
			childIndex++
		}
		if arr[parentIndex] <= arr[childIndex] {
			break
		}
		arr[childIndex], arr[parentIndex] = arr[parentIndex], arr[childIndex]
		parentIndex = childIndex
		childIndex = parentIndex*2 + 1
	}
}

// BuildHeap build a heap for array
func BuildHeap(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := len(arr) / 2; i >= 0; i-- {
		DownAdjust(arr, i, len(arr))
	}
}

// HeapSort is heap sort
func HeapSort(arr []int) {
	BuildHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		DownAdjust(arr, 0, i)
	}
}

func getParentIndex(childIndex int) int {
	if childIndex == 0 {
		return 0
	}
	if childIndex%2 == 0 {
		return childIndex/2 - 1
	}
	return childIndex / 2
}
