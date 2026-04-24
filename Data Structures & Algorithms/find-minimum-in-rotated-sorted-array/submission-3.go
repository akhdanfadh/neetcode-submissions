func findMin(nums []int) int {
	// intuition: O(logn) means dividing
	// observation: since sorted already, we can check where the min is
	//   by looking at each partition whether the start id and end id is still sorted
	
	sId, eId := 0, len(nums)-1
	for sId < eId {
		if nums[sId] <= nums[eId] { break } // already sorted
		mId := sId + (eId-sId)/2
		if nums[sId] > nums[mId] {
			eId = mId
		} else {
			sId = mId+1
		}
	}
	return nums[sId]
}
