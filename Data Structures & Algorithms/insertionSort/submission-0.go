// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	var res [][]Pair
	for i := 0; i < len(pairs); i++ {
		for j := i - 1; j >= 0; j-- {
			if pairs[j].Key > pairs[j+1].Key {
				pairs[j], pairs[j+1] = pairs[j+1], pairs[j]
			} else {
				break // left side is already sorted
			}
		}
		// in go, slices are reference, so need to create copy
		snapshot := make([]Pair, len(pairs))
		copy(snapshot, pairs)
		res = append(res, snapshot)
	}
	return res
}
