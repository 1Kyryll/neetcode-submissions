func isValid(s string) bool {
	if len(s)%2 != 0 {
        return false
    }

    for strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
        s = strings.ReplaceAll(s, "()", "")
        s = strings.ReplaceAll(s, "{}", "")
        s = strings.ReplaceAll(s, "[]", "")
    }   

    return s == ""
}
