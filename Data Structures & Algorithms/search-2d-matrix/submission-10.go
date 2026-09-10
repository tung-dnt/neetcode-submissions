func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	top,bot :=0,rows-1
	for top <= bot {
		mid := (top+bot)/2
		if matrix[mid][cols-1] < target {
			top = mid + 1
		} else if matrix[mid][0] > target {
			bot = mid - 1
		} else {
			break
		}
	}
	if !(top <= bot) {
        return false
    }

	row:=(top+bot)/2
	l,r:=0,len(matrix[row])-1
	for l <= r {
		mid := (l+r)/2
		if matrix[row][mid] < target {
			l = mid+1
		} else if matrix[row][mid] > target {
			r = mid-1
		} else {
			return true
		}
	}
	return false
}