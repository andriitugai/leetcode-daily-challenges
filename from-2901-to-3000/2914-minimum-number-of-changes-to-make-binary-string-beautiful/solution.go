func minChanges(s string) int {
    result := 0
    for i := 0; i < len(s); i += 2 {
        if s[i] != s[i + 1] {
            result += 1
        }
    }
    return result
}