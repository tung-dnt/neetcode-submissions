func simplifyPath(path string) string {
	words := strings.Split(path, "/")
	stack := []string{}
	for _, w := range words {
		if w == ".." {
			if len(stack) > 0 { stack = stack[:len(stack)-1] }
		} else if w != "" && w != "." {
			stack = append(stack, w)
		}
	}

	return "/" + strings.Join(stack, "/")
}