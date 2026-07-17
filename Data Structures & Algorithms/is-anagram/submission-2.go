func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap, tmap := make(map[byte]int), make(map[byte]int)

	for i:=0; i < len(s); i++ {
		smap[s[i]]++
		tmap[t[i]]++
	}
	for char := range smap {
		if smap[char] != tmap[char] {
			return false
		}
	}
	return true
}
