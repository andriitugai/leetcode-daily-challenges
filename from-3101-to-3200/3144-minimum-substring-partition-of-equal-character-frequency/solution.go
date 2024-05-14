func minimumSubstringsInPartition(s string) int {
    n := len(s)
    dp := make([]int, n + 1)
    for i := 1; i < n + 1; i ++ {
        dp[i] = n
    }

    for i := 0; i < n; i ++ {
        cnt := make(map[byte]int)
        for j := i; j < n; j ++ {
            cnt[s[j]] += 1
            if isBalanced(cnt) {
                dp[j + 1] = min(dp[j + 1], dp[i] + 1)
            }
        }
    }
    return dp[n]
}

func isBalanced(counts map[byte]int) bool {
    cVal := 0
    for _, v := range counts {
        if cVal == 0 {
            cVal = v
        } else {
            if v != cVal {
                return false
            }
        }
    }
    return true
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}