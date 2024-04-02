func countAlternatingSubarrays(nums []int) int64 {
    n := len(nums)
    dp := make([]int64, n)
    dp[0] = 1
    var result int64 = 1
    for i := 1; i < n; i ++ {
        dp[i] = 1
        if nums[i - 1] != nums[i] {
            dp[i] = dp[i - 1] + 1
        }
        result += dp[i]
    }
    return result
}