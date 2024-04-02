func minimumLevels(possible []int) int {
    n := len(possible)
    bp := 0
    for i := 0; i < n; i ++ {
        if possible[i] == 0 {
            possible[i] = -1
        }
        bp += possible[i]
    }
    dp := 0
    for lvl := 0; lvl < n - 1; lvl ++ {
        dp += possible[lvl]
        bp -= possible[lvl]
        if dp > bp {
            return lvl + 1
        }
    }
    return -1
}