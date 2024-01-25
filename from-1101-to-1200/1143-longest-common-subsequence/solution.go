func longestCommonSubsequence(text1 string, text2 string) int {
    n1, n2 := len(text1), len(text2)
    dp := make([][]int, n1 + 1)
    for i := 0; i <= n1; i ++ {
        dp[i] = make([]int, n2 + 1)
    }

    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }

    for i := n1 - 1; i >= 0; i -- {
        for j := n2 - 1; j >= 0; j -- {
            if text1[i] == text2[j] {
                dp[i][j] = 1 + dp[i + 1][j + 1]
            } else {
                dp[i][j] = max(dp[i + 1][j], dp[i][j + 1])
            }
        }
    }
    return dp[0][0]
}