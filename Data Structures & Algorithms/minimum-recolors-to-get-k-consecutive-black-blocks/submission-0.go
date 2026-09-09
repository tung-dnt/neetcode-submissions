func minimumRecolors(blocks string, k int) int {
	minOp,l,bCount := k,0,0
	for r:=0;r<len(blocks);r++{
		if blocks[r] == 'B' {
			bCount++
		}
		if r-l+1 > k {
			if blocks[l] == 'B' {
				bCount--
			}
			l++
		}
		if r-l+1 == k {
			minOp = min(minOp, k-bCount)
		}	
	}
	return minOp
}
