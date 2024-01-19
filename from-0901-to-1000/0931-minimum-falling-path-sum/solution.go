func minFallingPathSum(matrix [][]int) int {
    n := len(matrix)
    dp := make([][]int, n)
    for r := 0; r < n; r ++ {
        dp[r] = make([]int, n)
        copy(dp[r], matrix[r])
    }

    for r := 1; r < n; r ++ {
        for c := 0; c < n; c ++ {
            v := dp[r - 1][c]
            if c > 0 && dp[r - 1][c - 1] < v {
                v = dp[r - 1][c - 1]
            }
            if c < n - 1 && dp[r - 1][c + 1] < v {
                v = dp[r - 1][c + 1]
            }
            dp[r][c] += v
        }
    }

    result := dp[n - 1][0]
    for c := 1; c < n; c ++ {
        if result > dp[n-1][c] {
            result = dp[n - 1][c]
        }
    }
    return result
}