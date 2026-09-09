import "slices"

func sortStr(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}

func checkInclusion(s1, s2 string) bool {
	s1 = sortStr(s1)
	if len(s1) > len(s2) {
		return false
	}
	for i:=0; i+len(s1) <= len(s2); i++ {
		if s1 == sortStr(s2[i:i+len(s1)]) {
			return true
		}
	}
	return false
}