func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, p := range piles {
		r = max(r, p)
	}
	// Bất biến: đáp án luôn nằm trong [l, r]
	for l < r {
		mid := l + (r-l)/2 // mid lệch xuống
		if canFinish(piles, h, mid) {
			r = mid // mid thỏa → giữ mid làm ứng viên
		} else {
			l = mid + 1 // mid không thỏa → loại hẳn mid
		}
	}
	return l // l == r == tốc độ nhỏ nhất thỏa
}

func canFinish(piles []int, h, k int) bool {
	hours := 0
	for _, p := range piles {
		hours += (p + k - 1) / k
		if hours > h {
			return false
		}
	}
	return true
}