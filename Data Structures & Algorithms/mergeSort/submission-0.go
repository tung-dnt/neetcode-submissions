// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }
func merge(a, b []Pair) []Pair {
    res := make([]Pair, 0, len(a)+len(b))
    i, j := 0, 0
    for i < len(a) && j < len(b) {
        if a[i].Key <= b[j].Key {
            res = append(res, a[i]); i++
        } else {
            res = append(res, b[j]); j++
        }
    }
    res = append(res, a[i:]...)
    return append(res, b[j:]...)
}

func mergeSort(pairs []Pair) []Pair {
	if len(pairs) <= 1 {
		return pairs
	}
	mid := len(pairs) /2
	left := mergeSort(pairs[:mid])
	right := mergeSort(pairs[mid:])
	return merge(left, right)
}
