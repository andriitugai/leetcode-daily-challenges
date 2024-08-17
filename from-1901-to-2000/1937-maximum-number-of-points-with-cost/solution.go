func maxPoints(points [][]int) int64 {
    m, n := len(points), len(points[0])
    prev := make([]int64, n)
    for i := 0; i < n; i ++ {
        prev[i] = int64(points[0][i])
    }

    for i := 1; i < m; i ++ {
        leftMax := make([]int64, n)
        rightMax := make([]int64, n)

        leftMax[0] = prev[0]
        for j := 1; j < n; j ++ {
            leftMax[j] = max(leftMax[j - 1] - 1, prev[j])
        }

        rightMax[n - 1] = prev[n - 1]
        for j := n - 2; j >= 0; j -- {
            rightMax[j] = max(rightMax[j + 1] - 1, prev[j])
        }

        for j := 0; j < n; j ++ {
            prev[j] = int64(points[i][j]) + max(leftMax[j], rightMax[j])
        }
    }
    result := int64(prev[0])
    for i := 1; i < n; i ++ {
        result = max(result, prev[i])
    }
    return result
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}