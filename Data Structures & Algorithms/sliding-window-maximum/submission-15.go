func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    if n == 0 || k <= 0 || k > n {
        return nil
    }
    q := make([]int, 0, k)
    res := make([]int, 0, n-k+1)
    for i:=0;i<len(nums);i++ {
        if len(q)>0 && q[0] < i-k+1 {
            q = q[1:]
        }
        for len(q)>0 && nums[q[len(q)-1]] < nums[i] {
            q = q[:len(q)-1]
        }
        q = append(q, i)
        if i >= k-1 {
            res = append(res, nums[q[0]])
        }
    }
    return res
}