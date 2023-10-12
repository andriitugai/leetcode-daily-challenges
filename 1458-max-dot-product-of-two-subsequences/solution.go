unc maxDotProduct(nums1 []int, nums2 []int) int {
    m, n := len(nums1), len(nums2)
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dp[i][j] = nums1[i] * nums2[j]
            if i > 0 && j > 0 {
                dp[i][j] += max(dp[i-1][j-1], 0)
            }
            if i > 0 {
                dp[i][j] = max(dp[i][j], dp[i-1][j])
            }
            if j > 0 {
                dp[i][j] = max(dp[i][j], dp[i][j-1])
            }
        }
    }
    return dp[m-1][n-1]
}