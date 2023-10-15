func numWays(steps int, arrLen int) int {
    MOD := 1000000007
    m := steps + 1
    n := steps / 2 + 1
    if arrLen < n {
        n = arrLen
    }
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    dp[0][0] = 1
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            dp[i][j] = dp[i-1][j]
            if j > 0 {
                dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % MOD
            }
            if j < n - 1 {
                dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % MOD
            }
        }
    }
    return dp[steps][0] % MOD
}