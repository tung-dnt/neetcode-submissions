func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums), len(nums))
	
	prefix := 1
	for i, num := range nums {
		res[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res
}