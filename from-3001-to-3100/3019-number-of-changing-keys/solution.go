func countKeyChanges(s string) int {
    result := 0
    sl := strings.ToLower(s)
    for i := 1; i < len(s); i ++ {
        if sl[i] != sl[i-1] {
            result += 1
        }
    }
    return result
}