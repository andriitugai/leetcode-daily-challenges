func distributeCandies(n int, limit int) int {
    result := 0
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }
    for i := 0; i <= min(n, limit); i++ {
        for j := 0; j <= min(n - i, limit); j ++ {
            k := n - i - j
            if k >= 0 && k <= limit {
                result += 1
            }
        }
    }
    return result
}