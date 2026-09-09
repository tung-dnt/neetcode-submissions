func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    if n == 0 || k <= 0 || k > n {
        return nil
    }
    res := make([]int, 0, n-k+1)
    dq := make([]int, k)
    head, size := 0, 0

    for i := 0; i < n; i++ {
        if size > 0 && dq[head] <= i-k {
            head++
            if head == k {
                head = 0
            }
            size--
        }
        for size > 0 {
            t := head + size - 1
            if t >= k {
                t -= k
            }
            if nums[dq[t]] > nums[i] {
                break
            }
            size--
        }
        t := head + size
        if t >= k {
            t -= k
        }
        dq[t] = i
        size++

        if i >= k-1 {
            res = append(res, nums[dq[head]])
        }
    }
    return res
}