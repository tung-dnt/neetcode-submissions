// Need to revise
func characterReplacement(s string, k int) int {
	l,maxLen,maxF:=0,0,0
	freq := make(map[byte]int)
	for r:=0;r<len(s);r++ {
		freq[s[r]]++
		maxF = max(maxF, freq[s[r]])
		if (r-l+1)-maxF > k {
			freq[s[l]]--
			l++
		}
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
