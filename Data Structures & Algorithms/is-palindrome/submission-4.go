func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	left := 0
	right := len(s) - 1

	for left < right {
		rl := rune(s[left])
		if !unicode.IsLetter(rl) && !unicode.IsDigit(rl) {
			left++
			continue
		}

		rr := rune(s[right])
		if !unicode.IsLetter(rr) && !unicode.IsDigit(rr) {
			right--
			continue
		}

		if unicode.ToLower(rl) != unicode.ToLower(rr) {
			return false
		}

		left++
		right--
	}

	return true
}
