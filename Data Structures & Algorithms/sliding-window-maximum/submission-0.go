func maxSlidingWindow(nums []int, k int) []int {
	if k > len(nums) || k == 0 {
		return []int{}
	}

	res := []int{}

	for i := 0; i <= len(nums)-k; i++ {
		s := i + k

		cMax := max(nums[i:s])
		res = append(res, cMax)
	}

	return res
}

func max(nums []int) int {
	res := nums[0]

	for j := 1; j < len(nums); j++ {
		if res < nums[j] {
			res = nums[j]
		}
	}

	return res
}
