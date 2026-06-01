func checkInclusion(s1 string, s2 string) bool {
	// sliding window 
	if len(s1) > len(s2) {
        return false
    }

    var c1, c2 [26]int

    for i := 0; i < len(s1); i++ {
        c1[s1[i]-'a']++
        c2[s2[i]-'a']++
    }

    if c1 == c2 {
        return true
    }

    for i := len(s1); i < len(s2); i++ {    
        c2[s2[i]-'a']++
        c2[s2[i-len(s1)]-'a']--

        if c1 == c2 {
            return true
        }
    }

    return false
}
