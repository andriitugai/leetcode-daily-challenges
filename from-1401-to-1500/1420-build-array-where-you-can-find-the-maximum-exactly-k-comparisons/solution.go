func numOfArrays(n int, m int, k int) int {
    mod := 1000000007
    dp := make([][][]int, n + 1)
    for i := 0; i < n + 1; i++ {
        dp[i] = make([][]int, k + 1)
        for j := 0; j < k + 1; j ++ {
            dp[i][j] = make([]int, m + 1)
        }
    }
    for j := 1; j < m + 1; j++ {
        dp[1][1][j] = 1
    }

    for x := 1; x < n + 1; x++ {
        for y := 1; y < k + 1; y++ {
            for z := 1; z < m + 1; z++ {
                dp[x][y][z] += (dp[x-1][y][z] * z) % mod
                s := 0
                for i := 1; i < z; i++ {
                    s += (dp[x-1][y-1][i] % mod)
                }
                dp[x][y][z] += s
            }
        }
    }
    res := 0
    for i := 1; i < m + 1; i++ {
        res += (dp[n][k][i] % mod)
    }
    return res % mod
}