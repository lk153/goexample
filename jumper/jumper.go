package jumper

// You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
// Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:
// 0 <= j <= nums[i] and
// i + j < n
// Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

// Example 1:
// Input: nums = [2,3,1,1,4] Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Example 2:
// Input: nums = [2,3,0,1,4]
// Output: 2

// Constraints:
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// It's guaranteed that you can reach nums[n - 1].

// func main() {
// 	result := jumper.Exec([]int{2, 1, 1, 1, 4, 1, 1, 8})
// 	fmt.Println(result)
// }

func Exec(arr []int) int {
	maxIndex := len(arr) - 1
	totalStep := 0
	return exec(0, arr, maxIndex, totalStep)
}

func exec(idx int, arr []int, maxIndex, totalStep int) int {
	if idx >= maxIndex {
		return totalStep
	}

	maxStep := arr[idx]
	totalStep++
	min := -1
	for s := 1; s <= maxStep; s++ {
		nextIdx := idx + s
		tmp := exec(nextIdx, arr, maxIndex, totalStep)
		if min == -1 {
			min = tmp
			continue
		}

		if tmp <= min {
			min = tmp
		}
	}

	return min
}
