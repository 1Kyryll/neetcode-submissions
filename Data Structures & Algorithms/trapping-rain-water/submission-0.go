func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	maxLeft[0] = height[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	maxRight[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	water := 0
	for i := 0; i < n; i++ {
		water += min(maxLeft[i], maxRight[i]) - height[i]
	}

	return water
}
