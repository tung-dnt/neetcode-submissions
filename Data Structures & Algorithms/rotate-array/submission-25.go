func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}
	edge := n - k
	l, r := edge, n-1
	for l > 0 {
		iL := l - 1
		for i := 0; i < k; i++ {
			nums[iL+i], nums[iL+i+1] = nums[iL+i+1], nums[iL+i]
		}
		l--
		r--
	}
}

