func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		val := matrix[mid/n][mid%n]
		if val < target {
			left = mid + 1
		} else if val > target {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}