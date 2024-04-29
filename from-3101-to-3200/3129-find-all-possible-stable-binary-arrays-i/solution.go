func numberOfStableArrays(zero int, one int, limit int) int {
    MOD := 1000000007
    dp := make([][][]int, one + 1)
    for i := 0; i < one + 1; i ++ {
        dp[i] = make([][]int, zero + 1)
        for j := 0; j < zero + 1; j ++ {
            dp[i][j] = make([]int, 2)
        }
    }
    dp[0][0][0] = 1
    dp[0][0][1] = 1
    for i := 0; i < one + 1; i ++ {
        for j := 0; j < zero + 1; j ++ {
            for k := 1; k < limit + 1; k ++ {
                if i - k >= 0 {
                    dp[i][j][1] = (dp[i][j][1] + dp[i - k][j][0]) % MOD
                }
                if j - k >= 0 {
                    dp[i][j][0] = (dp[i][j][0] + dp[i][j - k][1]) % MOD
                }
            }
        }
    }
    return (dp[one][zero][0] + dp[one][zero][1]) % MOD
}