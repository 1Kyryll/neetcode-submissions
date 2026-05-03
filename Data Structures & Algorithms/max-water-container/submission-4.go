func maxArea(heights []int) int {
	max := 0
	left := 0
	right := len(heights) - 1

	for left < right {
		area := min(heights[left], heights[right]) * (right - left)
		if area > max {
			max = area
		}

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return max
}
