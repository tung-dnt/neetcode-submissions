func maxSlidingWindow(nums []int, k int) []int {
    var res []int
    var deque []int // stores indices of elements
    for i := 0; i < len(nums); i++ {
        // Remove elements outside current window
        if len(deque) > 0 && deque[0] < i-k+1 {
            deque = deque[1:]
        }
        // Remove smaller elements as they won't be the maximum
        for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
        if i >= k-1 {
            res = append(res, nums[deque[0]])
        }
    }
    return res
}