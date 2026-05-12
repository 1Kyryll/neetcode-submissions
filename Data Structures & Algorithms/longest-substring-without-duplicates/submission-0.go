func lengthOfLongestSubstring(s string) int {
	seen := map[byte]bool{}
	left, c := 0, 0

	for right := 0; right < len(s); right++ {
		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}

		seen[s[right]] = true

		if right-left+1 > c {
			c = right - left + 1
		}
	}

	return c
}
