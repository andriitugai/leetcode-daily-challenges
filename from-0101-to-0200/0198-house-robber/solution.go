func rob(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    var to_rob, not_to_rob int

    for i := 0; i < n; i ++ {
        to_rob = nums[i]
        if i > 1 {
            to_rob += dp[i - 2]
        }

        not_to_rob = 0
        if i > 0 {
            not_to_rob = dp[i - 1]
        }
        if not_to_rob > to_rob {
            dp[i] = not_to_rob
        } else {
            dp[i] = to_rob
        }
    }
    return dp[n - 1]
}