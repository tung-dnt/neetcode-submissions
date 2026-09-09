func simplifyPath(path string) string {
	stack := []string{}
	var subString strings.Builder

	for _, c := range path + "/" {
		s := subString.String()
		if c == '/' {
			if s == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else if s!="." && s!=""  {
				stack = append(stack, s)
			}
			subString.Reset()
		} else {
			subString.WriteRune(c)
		}
	}

	return "/" + strings.Join(stack, "/")
}