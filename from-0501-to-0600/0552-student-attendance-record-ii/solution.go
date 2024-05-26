func checkRecord(n int) int {
    mod := 1000000007
    dp := make([][]int, 2)
    for i := 0; i < 2; i ++ {
        dp[i] = make([]int, 3)
    }
    dp[0][0] = 1
    copyDP := func() [][]int {
        duplicate := make([][]int, 2)
        for i := 0; i < 2; i ++ {
            duplicate[i] = make([]int, 3)
            copy(duplicate[i], dp[i])
        }
        return duplicate
    }

    for i := 0; i < n; i ++ {
        prev := copyDP()
        dp[0][0] = (prev[0][0] + prev[0][1] + prev[0][2]) % mod
        dp[0][1] = prev[0][0]
        dp[0][2] = prev[0][1]
        dp[1][0] = (prev[0][0] + prev[0][1] + prev[0][2] + prev[1][0] + prev[1][1] + prev[1][2]) % mod
        dp[1][1] = prev[1][0]
        dp[1][2] = prev[1][1]
    }

    result := 0
    for r := 0; r < 2; r ++ {
        for c := 0; c < 3; c ++ {
            result = (result + dp[r][c]) % mod
        }
    }
    return result
}