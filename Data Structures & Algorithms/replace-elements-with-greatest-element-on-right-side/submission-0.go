func replaceElements(arr []int) []int {
	currentMax := -1
	
	for i:=len(arr)-1; i>=0; i-- {
		original := arr[i]
		arr[i] = currentMax
		currentMax = max(original, currentMax)
	}
	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}