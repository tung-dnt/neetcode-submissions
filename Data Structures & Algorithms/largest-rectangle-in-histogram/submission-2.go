type bar struct {
	start  int
	height int
}

/**
* Formula: h * i - start
* Point conditions:
* maxArea := 0
* if h > top.height: stack append({top.start, h})
* if h == top.height: stack append({?,h})
* if h < top.height: stack pop() -> max(maxArea, h), stack append({top.start, h})
 */

func largestRectangleArea(heights []int) int {
	var stack []bar
	maxArea := 0
	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1].height > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			start = top.start
			if area := top.height * (i - top.start); area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, bar{start, h})
	}

	n := len(heights)
	for _, item := range stack {
		if area := item.height * (n - item.start); area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}