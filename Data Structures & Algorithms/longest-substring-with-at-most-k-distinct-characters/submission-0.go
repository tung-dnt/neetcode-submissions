func lengthOfLongestSubstringKDistinct(s string, k int) int {
	l,maxLen := 0,0
	count:=make(map[byte]int)
	for r:=0;r<len(s);r++ {
		count[s[r]]++
		for len(count) > k {
			count[s[l]]--
			if count[s[l]] == 0 {
				delete(count, s[l])
			}
			l++
		}
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
