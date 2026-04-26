import "slices"

func groupAnagrams(strs []string) [][]string {
	sorted := make(map[string][]string)

	for k, v := range strs {
		sortedS := sortString(v)
		sorted[sortedS] = append(sorted[sortedS], strs[k])
	}

	result := [][]string{}

	for _, group := range sorted {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
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