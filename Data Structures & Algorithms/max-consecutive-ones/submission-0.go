func findMaxConsecutiveOnes(nums []int) int {
	maxCount, count := 0,0
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			maxCount = max(maxCount, count)
			count = 0
		}
	}	
	return max(maxCount, count)
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

