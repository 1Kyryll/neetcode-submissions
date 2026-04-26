import "slices"

func topKFrequent(nums []int, k int) []int {
	frequent := make(map[int]int)

	for _, v := range nums {
		if _, found := frequent[v]; found {
			frequent[v]++
		} else {
			frequent[v] = 1
		}
	}

	res := []int{}

	for k, _ := range frequent {
		res = append(res, k)
	}

	slices.SortFunc(res, func(a, b int) int {
		return frequent[b] - frequent[a] 
	})

	return res[:k]
}
