// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	res := make([][]Pair, 0)
	for i := 0; i < len(pairs); i++ {
		j := i - 1
		for j >= 0 && pairs[j+1].Key < pairs[j].Key {
			pairs[j], pairs[j+1] = pairs[j+1], pairs[j]
			j--
		}
		state := make([]Pair, len(pairs))
		copy(state, pairs)
		res = append(res, state)
	}
	return res
}