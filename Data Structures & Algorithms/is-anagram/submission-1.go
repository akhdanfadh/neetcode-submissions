func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[byte]int, len(s))
	for i := range len(s) {
		counts[s[i]]++
		counts[t[i]]--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}
