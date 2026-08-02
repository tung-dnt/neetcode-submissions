import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	for i:=0; i < len(nums); i++ {
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l, r := i+1, len(nums) -1
		for l < r {
			threeSum := a + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{nums[l], nums[r], a})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
