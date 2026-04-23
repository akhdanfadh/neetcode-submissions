func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		sorted := sortString(s)
		groups[sorted] = append(groups[sorted], s)
	}

	var res [][]string
	for _, g := range groups {
		res = append(res, g)
	}
	return res
}

func sortString(s string) string {
	var counts [26]int // input is lowercase english
	for i := range len(s) {
		counts[s[i]-'a']++
	}

	// rebuilt string
	res := []byte{}
	for i, c := range counts {
		for range c {
			res = append(res, byte(i)+'a')
		}
	}
	return string(res)
}