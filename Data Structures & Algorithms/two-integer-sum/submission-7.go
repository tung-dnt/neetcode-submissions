func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)
	for i,v := range nums {
		diff := target - v
		if j,found := prevMap[diff]; found {
			return []int{j, i}
		}
		prevMap[v] = i
	}

	return []int{}
}
