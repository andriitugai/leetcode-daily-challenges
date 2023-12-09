func kthDistinct(arr []string, k int) string {
    counts := make(map[string]int)
    for _, s := range arr {
        counts[s] += 1
    }
    c := 0
    for _, s := range arr {
        if counts[s] == 1 {
            c += 1
            if c == k {
                return s
            }
        }
    }
    return ""
}