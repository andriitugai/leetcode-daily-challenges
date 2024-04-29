func minEnd(n int, x int) int64 {
    res := int64(x)
    for i := 0; i < n - 1; i ++ {
        res = (res + 1) | int64(x)
    }
    return res
}