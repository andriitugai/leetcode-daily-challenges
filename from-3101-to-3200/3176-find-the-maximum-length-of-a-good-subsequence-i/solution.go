func maximumLength(nums []int, k int) int {
    dp := make([]map[int]int, k + 1)
    for i := 0; i <= k; i ++ {
        dp[i] = make(map[int]int)
    }
    result := make([]int, k + 1)
    for _, num := range nums {
        for i := k; i >= 0; i -- {
            if i == 0 {
                dp[0][num] = dp[0][num] + 1
            } else {
                dp[i][num] = max(dp[i][num] + 1, result[i - 1] + 1)
            }
            result[i] = max(result[i], dp[i][num])
        }
    }
    return result[k]
}