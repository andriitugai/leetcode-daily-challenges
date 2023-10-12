func integerBreak(n int) int {
    dp := make([]int, n+1)
    dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], dp[j]*(i-j))
			dp[i] = max(dp[i], j*(i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}