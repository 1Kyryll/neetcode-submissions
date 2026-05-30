func tokenize(text string, dictionary map[string]int) []string {
	output := []string{}
	pos := 0

	skeys := slices.Collect(maps.Keys(dictionary))
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(b), len(a))
	}

	slices.SortFunc(skeys, lenCmp)

	for pos < len(text) {
		matched := false
		for _, k := range skeys {
			if pos+len(k) <= len(text) {
				slice := text[pos : pos+len(k)]

				if slice == k {
					id := strconv.Itoa(dictionary[k])
					output = append(output, id)
					pos += len(k)
					matched = true
					break
				}
			}
		}

		if !matched {
			output = append(output, string(text[pos]))
			pos += 1
		}
	}

	return output
}