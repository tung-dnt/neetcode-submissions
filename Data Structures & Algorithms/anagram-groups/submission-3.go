func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[[26]int][]string)
	for _, str := range strs {
		// build map key
		var count [26]int
		for _, char := range str {
			count[char-'a']++
		}
		groupMap[count] = append(groupMap[count], str)
	}

	var result [][]string
	for _, group := range groupMap {
		result = append(result, group)
	}
	return result
}
