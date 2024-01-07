func numberOfArithmeticSlices(nums []int) int {
    n := len(nums)
    dp := make([]map[int]int, n)
    result := 0
    for i := 0; i < n; i ++ {
        dp[i] = make(map[int]int)
        for j := 0; j < i; j ++ {
            diff := nums[i] - nums[j]
            dp[i][diff] += (dp[j][diff] + 1)
            result += dp[j][diff]
        }
    }
    return result
}
