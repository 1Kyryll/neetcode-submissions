func maxProfit(prices []int) int {
	min := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		profit := prices[i] - min

		if profit > maxProfit {
			maxProfit = profit
		}
	}
	
	return maxProfit
}
