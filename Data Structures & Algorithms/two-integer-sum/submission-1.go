func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for k, v := range nums {
		complement := target - v

		if idx, found := seen[complement]; found {
			return []int{idx, k}
		}

		seen[v] = k
	}

	return nil
}
