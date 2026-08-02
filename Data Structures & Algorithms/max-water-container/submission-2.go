func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1
	maxArea := 0
	for l < r {
		area := (r-l) * min(heights[l], heights[r])
		if area > maxArea {
			maxArea = area
		}
		if heights[l] <= heights[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}