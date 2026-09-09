func minimumRecolors(blocks string, k int) int {
	minOp,l,whites := k,0,0
	for r:=0;r<len(blocks);r++{
		if blocks[r] == 'W' {
			whites++
		}
		if r-l+1 > k {
			if blocks[l] == 'W' {
				whites--
			}
			l++
		}
		if r-l+1 == k {
			minOp = min(minOp, whites)
		}
	}
	return minOp
}
