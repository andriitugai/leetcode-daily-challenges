func nthSuperUglyNumber(n int, primes []int) int {
    size := len(primes)
    dp := []int{1}
    ugly := 1
    idxs := make([]int, size)
    uglyNums := make([]int, size)
    for i := 0; i < size; i ++ {
        uglyNums[i] = 1
    }

    for i := 1; i < n; i ++ {
        for j := 0; j < size; j ++ {
            if uglyNums[j] == ugly {
                uglyNums[j] = dp[idxs[j]] * primes[j]
                idxs[j] += 1
            }
        }
        ugly = uglyNums[0]
        for _, u := range uglyNums[1:] {
            if u < ugly {
                ugly = u
            }
        }
        dp = append(dp, ugly)
    }
    return dp[n-1]
}