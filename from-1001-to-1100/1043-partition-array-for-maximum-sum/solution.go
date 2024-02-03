func maxSumAfterPartitioning(arr []int, k int) int {
    n := len(arr)
    dp := make([]int, n + 1)

    for start := n - 1; start >= 0; start -- {
        currMax := 0
        end := start + k
        if end > n {
            end = n
        }
        for i := start; i < end; i ++ {
            if arr[i] > currMax {
                currMax = arr[i]
            }
            alt := dp[i + 1] + currMax * (i - start + 1)
            if alt > dp[start] {
                dp[start] = alt
            }
        }
    }
    return dp[0]
}