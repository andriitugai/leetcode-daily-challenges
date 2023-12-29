func minDifficulty(jobDifficulty []int, d int) int {
    n := len(jobDifficulty)
    if n < d {
        return -1
    }
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, d + 1)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt / 2
        }
    }
    dp[0][0] = 0

    for i := 1; i <= n; i ++ {
        for k := 1; k <= d; k ++ {
            maxDiff := 0
            for j := i - 1; j >= k-1; j-- {
				if maxDiff < jobDifficulty[j] {
					maxDiff = jobDifficulty[j]
				}
				if dp[i][k] > dp[j][k - 1] + maxDiff {
					dp[i][k] = dp[j][k - 1] + maxDiff
				}
			}
        }
    }
    return dp[n][d]
}