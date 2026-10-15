// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }
type Solution struct{}

func NewSolution() *Solution { return &Solution{} }

func QuickSort(pairs []Pair) []Pair {
	quickSort(pairs)
	return pairs
}

func partition(pairs []Pair) int {
	l := -1
	pivot := len(pairs) - 1
	for r:=0; r < pivot; r++ {
		if pairs[r].Key < pairs[pivot].Key {
			l++
			pairs[l], pairs[r] = pairs[r], pairs[l]
		}
	}
	pairs[pivot], pairs[l+1] = pairs[l+1], pairs[pivot]
	return l + 1
}

func quickSort(pairs []Pair) {
	if len(pairs) <= 1 {
		return
	}
	pivot := partition(pairs)
	quickSort(pairs[:pivot])
	quickSort(pairs[pivot+1:])
}