// Need to revise
func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	maxf, l, res := 0, 0, 0
	for r := 0; r < len(s); r++ {
		count[s[r]]++
		maxf = max(count[s[r]], maxf)
		if (r-l+1)-maxf > k {
			count[s[l]]--
			l++
		}
		res = max(r-l+1, res)
	}
	return res
}