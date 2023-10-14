func paintWalls(cost []int, time []int) int {
    n := len(cost)
    inf := math.MaxInt32
    dp := make([]int, n + 1)
    for i := 1; i < n + 1; i ++ {
        dp[i] = inf
    }
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    for i := 0; i < n; i++ {
        for j := n; j > 0; j-- {
            dp[j] = min(dp[j], dp[max(j - time[i] - 1, 0)] + cost[i])
        }
    }
    return dp[n]
}