func productExceptSelf(nums []int) []int {
	result := []int{}

	for i := 0; i < len(nums); i++ {
		prod := 1
		for j, n := range nums {
			if j != i {
				prod *= n
			}
		}
		result = append(result, prod)
	}

	return result
}
