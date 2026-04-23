func hasDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, exists := set[v]; exists {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
