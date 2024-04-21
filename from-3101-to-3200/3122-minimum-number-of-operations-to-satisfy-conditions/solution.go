func minimumOperations(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([]int, 10)
    
    for col := n - 1; col >= 0; col -- {
        count := make([]int, 10)
        for i := 0; i < 10; i ++ {
            count[i] = m
        }
        for row := 0; row < m; row ++ {
            count[grid[row][col]] -= 1
        }

        minLeft := make([]int, 10)
        copy(minLeft, dp)
        minLeft[0] = 1000000000
        for i := 1; i < 10; i ++ {
            minLeft[i] = min(minLeft[i - 1], dp[i - 1])
        }

        minRight := make([]int, 10)
        copy(minRight, dp)
        minRight[9] = 1000000000
        for i := 8; i >= 0; i -- {
            minRight[i] = min(minRight[i + 1], dp[i + 1])
        }

        for i := 0; i < 10; i ++ {
            minLeft[i] = min(minLeft[i], minRight[i])
        }

        for i, ops := range count {
            dp[i] = ops + minLeft[i]
        }
    }
    result := dp[0]
    for i := 1; i < len(dp); i ++ {
        if dp[i] < result {
            result = dp[i]
        }
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}