const MOD = 1000000007

func kInversePairs(n int, k int) int {
    dp := make([][]int, n + 1)
    for i := 0; i < n + 1; i ++ {
        dp[i] = make([]int, k + 1)
    }

    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    dp[0][0] = 1
    for i := 1; i < n + 1; i ++ {
        for j := 0; j < k + 1; j ++ {
            for x := 0; x <= min(j, i - 1); x ++ {
                if j >= x {
                    dp[i][j] = (dp[i][j] + dp[i - 1][j - x]) % MOD
                }
            }
        }
    }
    return dp[n][k]
}