func findRotatedIdx(nums []int) int {
	l,r:=0,len(nums)-1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			return mid
		}
	}
	return 0
}

func binSearch(nums []int, target, l, r int) int {
	for l <= r {
		mid := (l+r)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	
	return -1
}

func search(nums []int, target int) int {
	rotated := findRotatedIdx(nums)
	firstFound := binSearch(nums, target, 0, rotated-1)
	if firstFound != -1 {
		return firstFound
	}
	return binSearch(nums, target, rotated, len(nums)-1)
}
