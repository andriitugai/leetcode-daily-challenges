func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    var n, m int = len(obstacleGrid), len(obstacleGrid[0])

    if obstacleGrid[0][0] == 1 || obstacleGrid[n-1][m-1] == 1 {
        return 0
    }

    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, m)
    }

    for r := 0; r < n; r ++ {
        for c := 0; c < m; c ++ {
            if obstacleGrid[r][c] == 1 {
                dp[r][c] = 0
            } else {
                if r == 0 && c == 0 {
                    dp[r][c] = 1
                } else if r == 0 && c > 0 {
                    dp[r][c] = dp[r][c - 1]
                } else if r > 0 && c == 0 {
                    dp[r][c] = dp[r-1][c]
                } else {
                    dp[r][c] = dp[r - 1][c] + dp[r][c - 1]
                }
            }
        }
    }

    return dp[n - 1][m - 1]
}