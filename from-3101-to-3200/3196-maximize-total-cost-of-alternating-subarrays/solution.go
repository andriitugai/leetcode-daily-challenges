func maximumTotalCost(nums []int) int64 {
    n := len(nums)
	dp := make([]int64, n+1)
	dp[1] = int64(nums[0])
	for i := 1; i < n; i++ {
		dp[i+1] = max(dp[i]+int64(nums[i]), dp[i-1]+int64(nums[i-1])-int64(nums[i]))
	}
	return dp[n]
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}