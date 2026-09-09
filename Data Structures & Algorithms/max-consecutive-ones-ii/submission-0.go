
func findMaxConsecutiveOnes(nums []int) int {
	l, maxLen := 0, 0
	nearest := -1
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			l = nearest + 1
			nearest = r
		}
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
