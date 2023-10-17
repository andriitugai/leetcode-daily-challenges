func isInterleave(s1 string, s2 string, s3 string) bool {
    m := len(s1)
    n := len(s2)

    if m + n != len(s3) {
        return false
    }

    dp := make([][]bool, m + 1)
    for i := 0; i < m + 1; i++ {
        dp[i] = make([]bool, n + 1)
    }
    dp[0][0] = true

    for i := 1; i < m + 1; i++ {
        dp[i][0] = (s1[i-1] == s3[i-1]) && dp[i-1][0]
    }

    for j := 1; j < n + 1; j ++ {
        dp[0][j] = (s2[j-1] == s3[j-1]) && dp[0][j-1]
    }

    for i := 1; i < m + 1; i++ {
        for j := 1; j < n + 1; j ++ {
            dp[i][j] = (s1[i-1] == s3[i+j-1] && dp[i-1][j]) || (s2[j-1] == s3[i+j-1] && dp[i][j-1])
        }
    }

    return dp[m][n]
}