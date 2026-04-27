type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
    for _, str := range strs {
        sb.WriteString(str)
        sb.WriteString("\x00")
    }
    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	parts := strings.Split(encoded, "\x00")
	return parts[:len(parts)-1]
}
