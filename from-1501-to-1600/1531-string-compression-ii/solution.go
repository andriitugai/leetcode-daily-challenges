func getLengthOfOptimalCompression(s string, k int) int {
    n := len(s)
    m := k

    dp := make([][]int, n + 1)
    for i := 0; i <= n; i ++ {
        dp[i] = make([]int, m + 1)
    }

    for i := 1; i <= n; i ++ {
        for j := 0; j <= i && j <= m; j ++ {
            need_remove := 0
            group_count := 0
            dp[i][j] = 101
            if j > 0 {
                dp[i][j] = dp[i - 1][j - 1]
            }

            for k := i; k >= 1; k -- {
                if s[k - 1] != s[i - 1] {
                    need_remove += 1
                } else {
                    group_count += 1
                }
            
                if need_remove > j {
                    break
                }

                mdp := dp[k - 1][j - need_remove]
                if group_count == 1 {
                    if dp[i][j] > mdp + 1 {
                        dp[i][j] = mdp + 1
                    }
                } else if group_count < 10 {
                    if dp[i][j] > mdp + 2 {
                        dp[i][j] = mdp + 2
                    }
                } else if group_count < 100 {
                    if dp[i][j] > mdp + 3 {
                        dp[i][j] = mdp + 3
                    }
                } else {
                    if dp[i][j] > mdp + 4 {
                        dp[i][j] = mdp + 4
                    }
                }
            }
        }
    }
    return dp[n][m]
}