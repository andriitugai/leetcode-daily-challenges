func minimumLength(s string) int {
    counts := make(map[byte]int)
    n := len(s)
    for i := 0; i < n; i ++ {
        counts[s[i]] += 1
    }
    toDelete := 0
    for _, v := range counts {
        if v % 2 == 0 {
            toDelete += v - 2
        } else {
            toDelete += v - 1
        }
    }
    return n - toDelete
}