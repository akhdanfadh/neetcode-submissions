func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[byte]int, len(s))
	mapT := make(map[byte]int, len(t))
	for i := range len(s) {
		mapS[s[i]]++
		mapT[t[i]]++
	}

	for c, vS := range mapS {
		vT, ok := mapT[c]
		if !ok || vS != vT {
			return false
		}
	}
	return true
}
