func replaceElements(arr []int) []int {
	currentMax := -1
	
	for i:=len(arr)-1; i>=0; i-- {
		original := arr[i]
		arr[i] = currentMax
		if original > currentMax {
			currentMax = original
		}
	}
	return arr
}