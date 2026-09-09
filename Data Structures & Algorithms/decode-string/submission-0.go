func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func decodeString(s string) string {
	stack := []string{}

	for _, c := range s {
		if c != ']' {
			stack = append(stack, string(c))
		} else {
			substr := ""
			for stack[len(stack)-1] != "[" {
				substr = stack[len(stack)-1] + substr
				stack = stack[:len(stack)-1]
			}

			stack = stack[:len(stack)-1] // remove '['

			freqStr := ""
			for len(stack) > 0 && isNumber(stack[len(stack)-1]) {
				freqStr = stack[len(stack)-1] + freqStr
				stack = stack[:len(stack)-1]
			}
			count, _ := strconv.Atoi(freqStr)
			stack = append(stack, strings.Repeat(substr, count))
		}
	}

	return strings.Join(stack, "")
}