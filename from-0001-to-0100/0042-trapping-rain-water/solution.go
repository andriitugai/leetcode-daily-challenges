func trap(height []int) int {
    n := len(height)
    if n == 0 {
        return 0
    }
    result := 0
    leftMax, rightMax := make([]int, n), make([]int, n)
    
    leftMax[0] = height[0]
    for i := 1; i < n; i ++ {
        leftMax[i] = max(height[i], leftMax[i - 1])
    }
    rightMax[n - 1] = height[n - 1]
    for i := n - 2; i >= 0; i -- {
        rightMax[i] = max(height[i], rightMax[i + 1])
    }

    for i := 1; i < n - 1; i ++ {
        result += min(leftMax[i], rightMax[i]) - height[i]
    }
    return result
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