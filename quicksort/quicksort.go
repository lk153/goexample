package quicksort

func partition(arr []int32, lowIdx, highIdx uint) ([]int32, uint) {
	pivot := arr[highIdx]
	i := lowIdx

	for j := lowIdx; j < highIdx; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[highIdx], arr[i] = arr[i], arr[highIdx]
	return arr, i
}

func QuickSort(arr []int32, lowIdx, highIdx uint) []int32 {
	if lowIdx > highIdx {
		return arr
	}

	arr, pi := partition(arr, lowIdx, highIdx)
	QuickSort(arr, lowIdx, pi-1)
	QuickSort(arr, pi+1, highIdx)

	return arr
}
