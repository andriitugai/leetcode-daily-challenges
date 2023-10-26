func numFactoredBinaryTrees(arr []int) int {
    MOD := 1000000007
    sort.Ints(arr)

    dp := make(map[int]int)
    for _, i := range arr {
        dp[i] = 1
    }

    for _, i := range arr {
        for _, j := range arr {
            if j * j > i {
                break
            }
            if i % j == 0 && dp[i / j] > 0 {
                tmp := (dp[j] * dp[i / j]) % MOD
                if j * j == i {
                    dp[i] = (dp[i] + tmp) % MOD
                } else {
                    dp[i] = (dp[i] + 2 * tmp) % MOD
                }
            }
        }
    }
    result := 0
    for _, val := range dp {
        result = (result + val) % MOD
    }
    return result
}