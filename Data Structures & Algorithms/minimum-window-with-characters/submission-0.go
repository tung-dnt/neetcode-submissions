func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	required := len(t)
	minLen := len(s) + 1
	bestL := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if tMap[s[r]] > 0 {
			required--
		}
		tMap[s[r]]--

		for required == 0 {
			if r-l+1 < minLen {
				minLen = r - l + 1
				bestL = l
			}
			tMap[s[l]]++
			if tMap[s[l]] > 0 {
				required++
			}
			l++
		}
	}

	if minLen > len(s) {
		return ""
	}
	return s[bestL : bestL+minLen]
}