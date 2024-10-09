package combinationtotal

// func main() {
// 	result := findCombination([]int32{2, 3, 6, 7}, 7)
// 	fmt.Println("RESULT:", result)
// }

func FindCombination(candidates []int32, target int32) (results [][]int32) {
	for _, i := range candidates {
		results = find(i, []int32{i}, candidates, target, results)
	}

	return
}

func find(total int32, combination []int32, candidates []int32, target int32, results [][]int32) [][]int32 {
	if total > target {
		return results
	}

	if total == target {
		quickSort(combination, 0, len(combination)-1)
		results = append(results, combination)
		return results
	}

	for _, i := range candidates {
		nextTotal := total + i
		nextCombination := append(combination, i)
		results = find(nextTotal, nextCombination, candidates, target, results)
	}

	return results
}

func quickSort(arr []int32, lowIdx int, highIdx int) []int32 {
	if lowIdx > highIdx {
		return arr
	}

	arr, pi := partition(arr, lowIdx, highIdx)
	quickSort(arr, lowIdx, pi-1)
	quickSort(arr, pi+1, highIdx)
	return arr
}

func partition(arr []int32, lowIdx int, highIdx int) ([]int32, int) {
	pivot := arr[highIdx]
	i := lowIdx

	for j := lowIdx; j < highIdx; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[highIdx] = arr[highIdx], arr[i]
	return arr, i
}
