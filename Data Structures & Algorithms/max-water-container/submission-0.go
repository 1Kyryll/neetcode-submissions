func maxArea(heights []int) int {
	max := 0 
	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			area := min(heights[i], heights[j]) * (j - i) 
			if area > max {
				max = area
			} 
		}
	}

	return max
}
