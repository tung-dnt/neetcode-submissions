/*
1. Find valid window by shift right pointer and update map needed
2. Once window is valid (required map keys count = 0) -> shrink from left
3. Update minLen as window shrink
4. If shrink to required map key, shift to the right and find if any char exist and repeat the process
*/
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	required:=len(t)
	minLen:=len(s)+1
	l,bestL:=0,0
	for r:=0; r<len(s); r++ {
		if count, exist := need[s[r]]; exist {
			if count > 0 {
				required--
			}
			need[s[r]]--
		}
		
		// Shrink from left when window is valid
		for required == 0 {
			if r-l+1 < minLen {
				minLen = r-l+1
				bestL=l
			}
			if _,exist := need[s[l]]; exist {
				need[s[l]]++
				if need[s[l]] > 0 {
					required++
				}				
			}
			l++
		}
	}
	if minLen > len(s) {
        return ""
    }
	return s[bestL:bestL+minLen]
}