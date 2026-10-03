
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}
	first, mid := 0, n-k
	next := mid
	for first != next {
		nums[first], nums[next] = nums[next], nums[first]
		first++
		next++
		if next == n {
			next = mid
		} else if first == mid {
			mid = next
		}
	}
}
