func longestIdealString(s string, k int) int {
    dp := make([]int, 256)
    for i := 0; i < len(s); i ++ {
        idx := int(s[i])
        dp[idx] = max(dp[idx - k: idx + k + 1]) + 1
    }
    return max(dp)
}

func max(arr []int) int {
    mx := arr[0]
    for _, x := range arr[1:] {
        if x > mx {
            mx = x
        }
    }
    return mx
}