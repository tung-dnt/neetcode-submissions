
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteByte('#')
		builder.WriteString(str)
	}

	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	left := 0
	for left < len(encoded) {
		right := left
		for encoded[right] != '#' {
			right++
		}
		length, _ := strconv.Atoi(encoded[left:right])
		left = right + 1
		res = append(res, encoded[left:left + length])
		left += length
	}
	return res
}