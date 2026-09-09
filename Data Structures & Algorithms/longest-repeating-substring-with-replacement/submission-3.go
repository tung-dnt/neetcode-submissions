func characterReplacement(s string, k int) int {
	l,maxLen,maxF:=0,0,0
	var freq [26]int
	for r:=0;r<len(s);r++ {
		freq[s[r]-'A']++
		maxF = max(maxF, freq[s[r]-'A'])
		if (r-l+1)-maxF > k {
			freq[s[l]-'A']--
			l++
		}
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
