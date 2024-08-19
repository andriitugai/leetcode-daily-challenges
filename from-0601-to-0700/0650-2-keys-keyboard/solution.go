func minSteps(n int) int {
    dp := make([]int, n + 1)
    for i := 0; i <= n; i ++ {
        dp[i] = 1000
    }
    dp[1] = 0

    for i := 2; i < n + 1; i ++ {
        for j := 1; j < i / 2 + 1; j ++ {
            if i % j == 0 {
                dp[i] = min(dp[i], dp[j] + i / j)
            }
        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}