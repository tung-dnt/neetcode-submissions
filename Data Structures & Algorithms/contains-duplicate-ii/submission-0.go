func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)
	for i, num := range nums {
		if j, ok := dict[num]; ok && i-j <= k {
			return true
		}
		dict[num] = i
	}
	return false
}
