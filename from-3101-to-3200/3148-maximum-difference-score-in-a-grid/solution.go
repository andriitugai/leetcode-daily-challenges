func maxScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m + 1)
    for i := 0; i < m + 1; i ++ {
        dp[i] = make([]int, n + 1)
    }

    result := math.MinInt32
    for i := m - 1; i >= 0; i -- {
        for j := n - 1; j >= 0; j -- {
            a, b, c, d := math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32
            if i < m - 1 {
                a = grid[i + 1][j] - grid[i][j] + dp[i + 1][j]
                c = grid[i + 1][j] - grid[i][j]
            }
            if j < n - 1 {
                b = grid[i][j + 1] - grid[i][j] + dp[i][j + 1]
                d = grid[i][j + 1] - grid[i][j]
            }

            dp[i][j] = max(a, max(b, max(c,d)))
            result = max(result, dp[i][j])
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}