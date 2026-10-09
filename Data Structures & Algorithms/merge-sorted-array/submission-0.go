func merge(nums1 []int, m int, nums2 []int, n int) {
	copied := make([]int, m)
	copy(copied, nums1[:m])
	i,j,k:= 0,0,0
	for i < m && j < n {
		if copied[i] <= nums2[j] {
			nums1[k] = copied[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	k += copy(nums1[k:], copied[i:])
	copy(nums1[k:], nums2[j:])
}