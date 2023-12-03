func kthFactor(n int, k int) int {
    kk := 0
    for i := 1; i <= n; i ++ {
        if n % i == 0 {
            kk += 1
            if kk == k {
                return i
            }
        }
    }
    return -1
}