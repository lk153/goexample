package twosum

func TwoSum(nums []int, target int) []int {
	for i1, v1 := range nums {
		remain := target - v1
		for i2, v2 := range nums {
			if v2 == remain && i1 != i2 {
				return []int{i1, i2}
			}
		}
	}

	return nil
}
