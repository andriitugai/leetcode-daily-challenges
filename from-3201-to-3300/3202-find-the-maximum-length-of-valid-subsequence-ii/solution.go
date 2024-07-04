func maximumLength(nums []int, k int) int {
    dp := make([][]int, k)
    for i := 0; i < k; i ++ {
        dp[i] = make([]int, k)
    }
    result := 1
    for _, num := range nums {
        currNum := num % k
        for mod := 0; mod < k; mod ++ {
            prevNum := (mod - currNum + k) % k
            dp[currNum][mod] = max(dp[currNum][mod], 1 + dp[prevNum][mod])
            result = max(result, dp[currNum][mod])
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}