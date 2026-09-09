func asteroidCollision(asteroids []int) []int {
	stack := []int{}
outer:
	for _, ast := range asteroids {
		for len(stack) > 0 && ast < 0 && stack[len(stack)-1] > 0 {
			diff := ast + stack[len(stack)-1]
			if diff > 0 { // top stack is larger 
				continue outer
			}
			stack = stack[:len(stack)-1]
			if diff == 0 {
				continue outer
			}
		}
		stack = append(stack, ast)
	}
	return stack
}