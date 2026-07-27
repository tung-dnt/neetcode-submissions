func longestConsecutive(nums []int) int {
	// array to set
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}
	
	// count longest consecutive
	longest := 0
	for num := range numSet {
		if _, ok := numSet[num - 1]; ok {
			continue
		}
	
		length := 1
		for {
			if _, ok := numSet[num + length]; !ok {
				break
			}
			length++
		}
			
		if longest < length {
			longest = length
		}
	}
	return longest
}
