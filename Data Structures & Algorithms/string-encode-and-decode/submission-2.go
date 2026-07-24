type Solution struct{}

/**
* To encode - add delimiter
* 1st decision: '#' as delimiter
* 	-> issue with test case ['tony', 'd#o']. Encode as 'tony#d#o' then Decode as [tony, d, o] -> incorrect
* 2nd decision: add number before delimiter as length
 */
func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteByte('#')
		builder.WriteString(str)
	}

	return builder.String()
}

/*
* 2 pointer:
* 	Test case: 5#Hello5#World
* */
func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0

	for i < len(encoded) {
		delimiterIdx := i
		for encoded[delimiterIdx] != '#' {
			delimiterIdx++
		}
		length, _ := strconv.Atoi(encoded[i:delimiterIdx])
		i = delimiterIdx + 1
		res = append(res, encoded[i:i+length])
		i += length
	}

	return res
}