unc isMatch(s string, p string) bool {
    m, n := len(p), len(s)
    dp := make([][]bool, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]bool, n+1)
    }

    for i := m; i >= 0; i-- {
        for j := n; j >= 0; j-- {
            if i == m && j == n {
                dp[i][j] = true
            } else if i == m {
                dp[i][j] = false
            } else if j == n {
                if p[i] == '*' {
                    dp[i][j] = dp[i+1][j]
                } else {
                    dp[i][j] = false
                }
            } else {
                if p[i] == '?' {
                    dp[i][j] = dp[i+1][j+1]
                } else if p[i] == '*' {
                    dp[i][j] = dp[i+1][j] || dp[i][j+1]
                } else if p[i] == s[j] {
                    dp[i][j] = dp[i+1][j+1]
                } else {
                    dp[i][j] = false
                }
            }
        }
    }
    return dp[0][0]
}