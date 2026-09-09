func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	l, alreadySatisfied, window, maxWindow := 0,0,0,0
	for r:=0; r<len(customers); r++ {
		if grumpy[r] == 1 {
			window += customers[r]
		} else {
			alreadySatisfied+=customers[r]
		}
		if r-l+1 > minutes {
			if grumpy[l] == 1 {
				window-=customers[l]
			}
			l++
		}
		maxWindow = max(window, maxWindow)
	}

	return maxWindow + alreadySatisfied
}
