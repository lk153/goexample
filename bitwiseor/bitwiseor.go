package bitwiseor

func FindLongestSlice(arr []int) []int {
	min := -1
	minIdx := -1
	for idx, item := range arr {
		if min == -1 {
			min = item
			minIdx = idx
			continue
		}

		if min > item {
			min = item
			minIdx = idx
		}
	}

	maxIdx := minIdx
	flagCount := false
	prevIdxMax := minIdx
	prevIdxMin := minIdx

	for idx, item := range arr {
		switch {
		case !flagCount && item != min:
		case !flagCount && item == min:
			prevIdxMin = idx
			prevIdxMax = idx
			flagCount = true
		case flagCount && item == min:
			prevIdxMax = idx
		case flagCount && item != min:
			if (prevIdxMax - prevIdxMin) > (maxIdx - minIdx) {
				maxIdx = prevIdxMax
				minIdx = prevIdxMin
			}
			flagCount = false
		}
	}

	if (prevIdxMax - prevIdxMin) > (maxIdx - minIdx) {
		maxIdx = prevIdxMax
		minIdx = prevIdxMin
	}

	return arr[minIdx : maxIdx+1]
}
