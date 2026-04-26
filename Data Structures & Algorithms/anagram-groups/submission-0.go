func groupAnagrams(strs []string) [][]string {
	matrix := [][]string{}
	visited := make([]bool, len(strs))

	for i := 0; i < len(strs); i++ {
		if visited[i] {
			continue
		}

		group := []string{strs[i]}

		for j := i + 1; j < len(strs); j++ {
			if !visited[j] && isAnagram(strs[i], strs[j]) {
				group = append(group, strs[j])
				visited[j] = true
			}
		}

		visited[i] = true
		matrix = append(matrix, group)
	}

	return matrix
}

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