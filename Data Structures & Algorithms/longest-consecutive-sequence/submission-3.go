import "slices"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)
	res := 1
	maxRes := 1

	for i := 0; i < len(nums); i++ {
		if res > maxRes {
			maxRes = res
		}
		if i+1 < len(nums) {
			if nums[i]+1 == nums[i+1] {
				res++
			} else if nums[i] == nums[i+1] {
				continue
			} else {
				res = 1
			}
		}
	}

	return maxRes
}
