func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if r == 0 || c == 0 {
                dp[r][c] = 1
            } else {
                dp[r][c] = dp[r-1][c] + dp[r][c-1]
            }
        }
    }
    return dp[m-1][n-1]