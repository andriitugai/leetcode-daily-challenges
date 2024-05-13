func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    dp := make([]int, n)
    for i := n - 1; i >= 0; i -- {
        dp[i] = energy[i]
        if i + k < n {
            dp[i] += dp[i + k]
        }
    }
    result := dp[0]
    for i := 1; i < n; i ++ {
        if dp[i] > result {
            result = dp[i]
        }
    }
    return result
}