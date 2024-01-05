func lengthOfLIS(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }

    dp := make([]int, n)
    for i := 0; i < n; i ++ {
        dp[i] = 1
    }

    result := 1
    for i := 0; i < n; i ++ {
        for j := 0; j < i; j ++ {
            if nums[i] > nums[j] && dp[j] + 1 > dp [i] {
                dp[i] = dp[j] + 1
            }
        }
        if dp[i] > result {
            result = dp[i]
        }
    }
    return result
}