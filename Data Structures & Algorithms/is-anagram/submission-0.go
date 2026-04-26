func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[rune]int)

	for i, ch := range s {
		freq[ch]++
		freq[rune(t[i])]--
	}

	for _, count := range freq {
		if count != 0 {
			return false
		}
	}

	return true
}


