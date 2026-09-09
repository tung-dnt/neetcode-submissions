func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)
	for R, num := range nums {
		if L, ok := dict[num]; ok && R-L <= k {
			return true
		}
		dict[num] = R
	}
	return false
}
