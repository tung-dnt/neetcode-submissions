// Need to revise
func characterReplacement(s string, k int) int {
	var freq [26]int
	maxf, l, res := 0, 0, 0
	for r := 0; r < len(s); r++ {
		freq[s[r]-'A']++
		maxf = max(freq[s[r]-'A'], maxf)
		if (r-l+1)-maxf > k {
			freq[s[l]-'A']--
			l++
		}
		res = max(r-l+1, res)
	}
	return res
}
