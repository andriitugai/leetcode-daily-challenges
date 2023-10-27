func longestPalindrome(s string) string {
    if len(s) == 1 {
        return s
    }
    n := len(s)
    maxPal := s[0:1]

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        dp[i][i] = 1
    }

    for i := 0; i < n-1; i++ {
        if s[i] == s[i+1] {
            dp[i][i+1] = 1
            maxPal = s[i:i+2]
        }
    }

    for offset := 2; offset < n; offset ++ {
        for i := 0; i + offset < n; i++ {
            if s[i] == s[i + offset] && dp[i + 1][i + offset - 1] == 1 {
                dp[i][i + offset] = 1
                if len(maxPal) < offset + 1 {
                    maxPal = s[i:i+offset+1]
                }
            }
        }
    }
    return maxPal
}
