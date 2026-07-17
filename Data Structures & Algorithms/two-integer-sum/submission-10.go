func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)
	for i,v := range nums {
		if j,found := prevMap[target - v]; found {
			return []int{j, i}
		} else {
			prevMap[v] = i
		}
	}

	return []int{0, 0}
}
