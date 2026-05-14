func characterReplacement(s string, k int) int {
	var count [26]int
	left, result, maxFreq := 0, 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++
		if count[s[right]-'A'] > maxFreq {
			maxFreq = count[s[right]-'A']
		}

		if (right-left+1)-maxFreq > k {
			count[s[left]-'A']--
			left++
		}

		if right-left+1 > result {
			result = right - left + 1
		}
	}

	return result
}
