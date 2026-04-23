func groupAnagrams(strs []string) [][]string {
	mapping := make(map[int][]int)
	mapped := make(map[int]bool, len(strs))

	for i := 0; i < len(strs); i++ {
		if mapped[i] { continue }
		mapped[i] = true
		mapping[i] = []int{}
		for j := i + 1; j < len(strs); j++ {
			if isAnagram(strs[i], strs[j]) {
				mapping[i] = append(mapping[i], j)
				mapped[j] = true
			}
		}
	}

	result := [][]string{}
	for k, v := range mapping {
		res := []string{strs[k]}
		for _, s := range v {
			res = append(res, strs[s])
		}
		result = append(result, res)
	}
	return result
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [26]int // input is lowercase english
	for i := range len(s) {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}