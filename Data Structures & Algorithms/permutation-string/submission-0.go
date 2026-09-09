import "slices"

func checkInclusion(s1, s2 string) bool {
	n := len(s1)
	if n > len(s2) {
		return false
	}
	target := sortString(strings.ToLower(s1))
	for l := 0; l+n <= len(s2); l++ {
		if target == sortString(strings.ToLower(s2[l:l+n])) {
			return true
		}
	}
	return false
}

func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}
