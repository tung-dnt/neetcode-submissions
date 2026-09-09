
func numOfSubarrays(arr []int, k int, threshold int) int {
	threshold*=k
	count:=0
	for i:=0; i+k<=len(arr);i++ {
		sum:=0
		for _, num := range arr[i:i+k] {
			sum+=num
		}
		if sum >= threshold {
			count++
		}
	}
	return count
}