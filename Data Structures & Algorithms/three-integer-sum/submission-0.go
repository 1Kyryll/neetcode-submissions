import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return res
}
