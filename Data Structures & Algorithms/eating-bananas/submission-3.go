/*
totalTime <= h -> mid large -> mid- -> r- (loop until found minumum eligible)
totalTime > h -> mid small -> mid+ -> l+
*/
func minEatingSpeed(piles []int, h int) int {
	l,r:=1,0
	for _,p:=range piles {
		r=max(r,p)
	}
	res:=r

	for l <= r {
		mid:= l + (r-l)/2
		totalTime:=0
		for _, p:=range piles {
			totalTime+=int(math.Ceil(float64(p)/float64(mid)))
		}
		if totalTime <= h {
			res = mid
			r = mid -1
		} else {
			l = mid +1
		}
	}

	return res
}
