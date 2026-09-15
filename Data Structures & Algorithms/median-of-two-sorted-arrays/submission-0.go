import "slices"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := append(nums1, nums2...)

	slices.Sort(arr)
	
	if len(arr) % 2 == 0 {
		left := arr[len(arr)/2-1]
		right := arr[len(arr)/2]
		return float64(left+right) / 2.0
	}

	return float64(arr[len(arr)/2])
}
