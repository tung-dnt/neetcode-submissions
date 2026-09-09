func maxProfit(prices []int) int {
	best := 0
	l := 0
	for r:=1; r < len(prices); r++ {
		if prices[l] < prices[r] {
			best = max(best, prices[r] - prices[l])
		} else {
			l = r
		}
	}
	return best
}
