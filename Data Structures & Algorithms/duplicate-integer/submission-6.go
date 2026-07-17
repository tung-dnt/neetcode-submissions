func hasDuplicate(nums []int) bool {
	counts := make(map[int]struct{})
	for _, num := range nums {
		counts[num] = struct{}{}
	}
	return len(counts) < len(nums)
}
