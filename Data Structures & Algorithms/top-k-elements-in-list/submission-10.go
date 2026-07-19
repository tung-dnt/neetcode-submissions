// import (
// 	"slices"
// 	"cmp"
// )
// func topKFrequent(nums []int, k int) []int {
// 	// 1. Map element frequencies
// 	counts := make(map[int]int)
// 	for _, num := range nums {
// 		counts[num]++
// 	}

// 	// 2. Define a local 2-tuple struct and collect map entries
// 	type entry struct {
// 		count int
// 		num int
// 	}
// 	arr := make([]entry, 0, len(counts))
// 	for num, count := range counts {
// 		arr = append(arr, entry{count, num})
// 	}
// 	// 3. Sort keys
// 	slices.SortFunc(arr, func (a, b entry) int {
// 		return cmp.Compare(b.count, a.count)
// 	})

// 	// 4. Extract the top K elements
// 	res := make([]int, k)
// 	for i := range k {
// 		res[i] = arr[i].num
// 	}
	
// 	return res
// }

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	bucket := make([][]int, len(nums)+1)
	for num, count := range counts {
		bucket[count] = append(bucket[count], num)
	}

	result := make([]int, 0, k)
	for i := len(bucket) - 1; i > 0 && len(result) < k; i-- {
		for _, item := range bucket[i] {
			result = append(result, item)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}