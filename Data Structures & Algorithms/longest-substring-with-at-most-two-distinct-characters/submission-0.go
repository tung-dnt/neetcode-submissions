func lengthOfLongestSubstringTwoDistinct(s string) int {
    freq := make(map[byte]int)
	maxLen,l := 0,0
	for r:=0; r<len(s); r++ {
		freq[s[r]]++
		for len(freq) > 2 {
			if freq[s[l]]==1 {
				delete(freq,s[l])
			} else {
				freq[s[l]]--
			}
			l++
		}
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
