func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    if n == 0 || k <= 0 || k > n {
        return nil
    }
    res := make([]int, 0, n-k+1)
    dq := make([]int, k) // đúng k, không hơn
    head, size := 0, 0

    for i := 0; i < n; i++ {
        if size > 0 && dq[head] <= i-k {
            head = (head + 1) % k
            size--
        }
        for size > 0 && nums[dq[(head+size-1)%k]] <= nums[i] {
            size--
        }
        dq[(head+size)%k] = i
        size++
        if i >= k-1 {
            res = append(res, nums[dq[head]])
        }
    }
    return res
}