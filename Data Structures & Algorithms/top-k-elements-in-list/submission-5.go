import (
  "slices"
  "cmp"
)
func topKFrequent(nums []int, k int) []int {
	// 1. Map element frequencies
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// 2. Define a local 2-tuple struct and collect map entries
	type entry struct {
		num, cnt int
	}
	
	arr := make([]entry, 0, len(count))
	for num, cnt := range count {
		arr = append(arr, entry{num, cnt})
	}

	// 3. Sort the entries in descending order based on count
	slices.SortFunc(arr, func(a, b entry) int {
		return cmp.Compare(b.cnt, a.cnt)
	})

	// 4. Extract the top K elements
	res := make([]int, k)
	for i := range k {
		res[i] = arr[i].num
	}
	
	return res
}