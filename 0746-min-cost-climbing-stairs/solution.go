func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost) + 2)
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }
    for i := len(cost) - 1; i >= 0; i-- {
        dp[i] = cost[i] + min(dp[i+1], dp[i+2])
    }
    return min(dp[0], dp[1])
}