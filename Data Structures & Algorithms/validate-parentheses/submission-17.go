func isValid(s string) bool {
    lookup := map[string]string{
		"]": "[",
		"}": "{",
		")": "(",
	}
	stack := make([]string, 0)
	for _, c := range s {
		if open, exists := lookup[string(c)]; exists {
			if len(stack) == 0 {
				return false
			} else {
				top := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				if top != open {
					return false
				}
			}
		} else {
			stack = append(stack, string(c))
		}
	}
	return len(stack) == 0
}
