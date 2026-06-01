func checkInclusion(s1 string, s2 string) bool {
	// brute force
	if len(s1) > len(s2) {
		return false
	}

	seen1 := make(map[rune]int)

	for _, r := range s1 {
		seen1[r]++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		subseen := make(map[rune]int)
		sub := s2[i : i+len(s1)]
		for _, r := range sub {
			subseen[r]++
		}

		match := true
		for r, c := range seen1 {
			if subseen[r] != c {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}
