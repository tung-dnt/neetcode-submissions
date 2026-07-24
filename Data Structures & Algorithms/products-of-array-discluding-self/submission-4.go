func productExceptSelf(nums []int) []int {
	product := 1
	zeroCount := 0
	res := make([]int, len(nums), len(nums))

	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			product *= num
		}
	}
	if zeroCount > 1 {
		return res
	}
	for i, num := range nums {
		if zeroCount > 0 {
			if num == 0 {
				res[i] = product
			} else {
				res[i] = 0
			}
		} else {
			res[i] = product / num
		}
	}

	return res
}