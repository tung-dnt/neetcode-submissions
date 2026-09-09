
func numOfSubarrays(arr []int, k int, threshold int) int {
	threshold*=k
	l,count,sum:=0,0,0
	for r:=0; r<len(arr);r++ {
		sum+=arr[r]
		if r-l+1 >= k {
			if sum >= threshold {
				count++
			}
			sum-=arr[l]
			l++
		}
	}
	return count
}