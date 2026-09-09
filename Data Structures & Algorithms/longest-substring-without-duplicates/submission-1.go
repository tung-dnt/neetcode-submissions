func lengthOfLongestSubstring(s string) int {
	l, maxLen := 0, 0
	dups := make(map[rune]int)
	for r, char := range s {
		if pos, ok := dups[char]; ok && pos >= l {
			l = pos + 1
		}
		dups[char] = r
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}