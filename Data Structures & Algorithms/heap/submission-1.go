type MinHeap struct {
	arr []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func getParId(i int) int { return (i - 1) / 2 }
func getLeftId(i int) int { return 2 * i + 1 }
func getRightId(i int) int { return 2 * (i + 1) }

func (mh *MinHeap) Push(val int) {
	mh.arr = append(mh.arr, val)
	// childs should always be >= parent
	iNew := len(mh.arr) - 1
	for iNew > 0 {
		iPar := getParId(iNew)
		if mh.arr[iNew] >= mh.arr[iPar] {
			break
		} 
		mh.arr[iNew], mh.arr[iPar] = mh.arr[iPar], mh.arr[iNew]
		iNew = iPar
	}
}

func (mh *MinHeap) percolateDown(iPar int) {
	length := len(mh.arr)
	for length > getLeftId(iPar) { // guard agaings out of bound
		iLeft, iRight, iSmall := getLeftId(iPar), getRightId(iPar), iPar

		// check if this Par should be swap with Left or Right
		if length > iLeft && mh.arr[iLeft] < mh.arr[iSmall] { iSmall = iLeft }
		if length > iRight && mh.arr[iRight] < mh.arr[iSmall] { iSmall = iRight }
		if iPar == iSmall { break } // if stay the same, nothing to percolate

		// at this point, it is guaranteed that iSmall changes, so swap
		mh.arr[iPar], mh.arr[iSmall] = mh.arr[iSmall], mh.arr[iPar]
		iPar = iSmall
	}
}

func (mh *MinHeap) Pop() int {
	if len(mh.arr) == 0 { return -1 }
	
	val := mh.arr[0]
	if len(mh.arr) == 1 {
		mh.arr = []int{}
		return val
	}

	// mh.arr = mh.arr[1:] // this shifts the array, so unnecessary O(n)
	mh.arr[0] = mh.arr[len(mh.arr)-1] // swap root with last and remove last
	mh.arr = mh.arr[:len(mh.arr)-1]
	mh.percolateDown(0)
	return val
}

func (mh *MinHeap) Top() int {
	if len(mh.arr) == 0 { return -1 }
	return mh.arr[0]
}

func (mh *MinHeap) Heapify(nums []int) {
	// naive approach would be push every element, but that is O(nlogn)

	// building the intution:
	// 1. what does a valid heap need?
	// 	  all parents need to be >= children
	// 2. can i fix one node at a time?
	//    well, if we pick any node that violates this and swap it with its smallest,
	//    that child might violate the property above. that's why the percolateDown
	// 3. what order do i process node then?
	//    - top to bottom -> problem: when fix a node by swapping down, might break
	//      node that already fixed above. percolateDown assumes subtrees already valid
	//    - bottom to top -> leaf has no children so it's trivially valid. its parent
	//      can now safely swap down knowing the subtree below is already valid
	// 4. then how can we get the leaf?
	//    notice that half the array is the leaf

	mh.arr = nums
	for iPar := (len(nums) - 1) / 2; iPar >= 0; iPar-- {
		// fix all non-leaf node starting from the last non-leaf node
 		mh.percolateDown(iPar) 
	}
}
