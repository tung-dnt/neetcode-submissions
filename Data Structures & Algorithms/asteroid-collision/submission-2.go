func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, ast := range asteroids {
		// When collision happens
		for len(stack) > 0 && ast < 0 && stack[len(stack)-1] > 0 {
			diff := stack[len(stack)-1] + ast
			if diff < 0 {
				stack = stack[:len(stack)-1]
			} else if diff > 0 {
				ast = 0
			} else if diff == 0 {
				stack = stack[:len(stack)-1]
				ast = 0
			}
		}
		if ast != 0 {
			stack = append(stack, ast)
		}
	}
	return stack
}