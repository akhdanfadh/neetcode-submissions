func topKFrequent(nums []int, k int) []int {
	// count frequency
	byNum := make(map[int]int, len(nums))
	for _, num := range nums {
		byNum[num]++
	}

	// bucket sort
	byCount := make([][]int, len(nums)+1) // this +1 is big
	for num, count := range byNum {
		byCount[count] = append(byCount[count], num)
	}

	// get top k
	var res []int
	for i := len(byCount)-1; i > 0; i-- {
		if k == 0 { break }
		for j := range byCount[i] {
			if k == 0 { break }
			res = append(res, byCount[i][j])
			k--
		}
	} 
	return res
}
