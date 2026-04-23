// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func QuickSort(pairs []Pair) []Pair {
	qSort(pairs, 0, len(pairs)-1)
	return pairs
}

func qSort(pairs []Pair, iStart, iEnd int) {
	if iStart >= iEnd { return } // base: 0/1 element

	iPivot := partition(pairs, iStart, iEnd)
	qSort(pairs, iStart, iPivot-1)
	qSort(pairs, iPivot+1, iEnd)
}

func partition(pairs []Pair, iStart, iEnd int) int {
	pivot := pairs[iEnd].Key
	iBound := iStart // less than pivot boundary
	for iScan := iStart; iScan < iEnd; iScan++ {
		if pairs[iScan].Key < pivot {
			pairs[iBound], pairs[iScan] = pairs[iScan], pairs[iBound]
			iBound++
		}
	}
	// place pivot in its final
	pairs[iBound], pairs[iEnd] = pairs[iEnd], pairs[iBound]
	return iBound
}