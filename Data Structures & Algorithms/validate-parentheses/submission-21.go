func isValid(s string) bool {
    lookup := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := make([]byte, 0)
	for i:=0; i< len(s); i++ {
		if open, exists := lookup[s[i]]; exists {
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
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
