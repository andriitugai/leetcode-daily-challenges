func valueAfterKSeconds(n int, k int) int {
    a := make([]int, n)
    for i := 0; i < n; i ++ {
        a[i] = 1
    }

    mod := 1000000007
    for j := 0; j < k; j ++ {
        for i := 1; i < n; i ++ {
            a[i] = (a[i - 1] + a[i]) % mod
        }
    }
    return a[n - 1]
}