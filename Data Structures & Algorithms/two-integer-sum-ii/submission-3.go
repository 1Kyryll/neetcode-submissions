func twoSum(numbers []int, target int) []int {
	seen := make(map[int]int) 
	
	for i, num := range numbers {
		complement := target - num 
		
		if j, ok := seen[complement]; ok {
			return []int{j+1, i+1}
		}
		
		seen[num] = i
	}
	
	return nil 
}
