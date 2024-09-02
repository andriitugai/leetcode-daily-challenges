func chalkReplacer(chalk []int, k int) int {
    n := len(chalk)
    s := 0
    for i := 0; i < n; i ++ {
        s += chalk[i]
    }
    k %= s
    for i := 0; i < n; i ++ {
        if chalk[i] > k {
            return i
        }
        k -= chalk[i]
    }
    return -1
}