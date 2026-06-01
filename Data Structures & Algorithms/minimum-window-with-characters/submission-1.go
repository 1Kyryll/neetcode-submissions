func minWindow(s string, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}
	
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	
	missing := len(t) 
	left := 0 
	start, end := 0,0 
	
	for right := 0; right < len(s); right++ {
		if need[s[right]] > 0 {
			missing-- 
		}
		need[s[right]]--
		
		for missing == 0 {
			if end == 0 || right-left+1 < end-start {
				start, end = left, right+1
			}
			
			need[s[left]]++
			if need[s[left]] > 0 {
				missing++
			}
			left++
		}
	}
	
	return s[start:end]
}
