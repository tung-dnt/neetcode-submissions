func trap(height []int) int {
    n := len(height)
    if n == 0 {
        return 0
    }

    leftMax := make([]int, n)
    rightMax := make([]int, n)

    // Initialize the boundaries.
    leftMax[0] = height[0]
    rightMax[n-1] = height[n-1]

    // Fill both arrays in one loop:
    // leftMax from left to right
    // rightMax from right to left
    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i])

        rightIdx := n - 1 - i
        rightMax[rightIdx] = max(
            rightMax[rightIdx+1],
            height[rightIdx],
        )
    }

    water := 0
    for i := 0; i < n; i++ {
        water += min(leftMax[i], rightMax[i]) - height[i]
    }

    return water
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}