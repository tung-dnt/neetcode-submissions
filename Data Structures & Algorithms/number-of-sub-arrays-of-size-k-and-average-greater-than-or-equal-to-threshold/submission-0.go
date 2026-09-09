func numOfSubarrays(arr []int, k int, threshold int) int {
	curSum, res, l := 0, 0, 0
	target := k * threshold
	for r, num := range arr {
		curSum += num
		if r-l+1 == k {
			if curSum >= target {
				res++
			}
			curSum -= arr[l]
			l++
		}
	}
	return res
}