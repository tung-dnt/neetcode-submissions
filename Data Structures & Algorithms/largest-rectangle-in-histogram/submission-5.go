type bar struct {
	start  int
	height int
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([]bar, 0, len(heights))

	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1].height > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxArea = max(top.height*(i-top.start), maxArea)
			start = top.start
		}
		stack = append(stack, bar{start: start, height: h})
	}

	for _, b := range stack {
		maxArea = max(b.height*(len(heights)-b.start), maxArea)
	}

	return maxArea
}