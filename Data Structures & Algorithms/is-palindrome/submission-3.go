func isPalindrome(s string) bool {
	// brute force solution

	trimmedS := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			return -1
		}
		return r
	}, s)
	fmt.Println(trimmedS)
	sReversedChars := make([]string, 0, len(s))

	for i := len(trimmedS) - 1; i >= 0; i-- {
		sReversedChars = append(sReversedChars, string(trimmedS[i]))
	}

	sReversed := strings.Trim(strings.Join(sReversedChars, ""), "?!.,'`")
	fmt.Println(sReversed)

	return strings.ToLower(sReversed) == strings.ToLower(trimmedS)
}
