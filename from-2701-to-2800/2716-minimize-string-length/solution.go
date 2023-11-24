func minimizedStringLength(s string) int {
    c := make(map[byte]bool)
    for i := 0; i < len(s); i ++ {
        c[s[i]] = true
    }
    return len(c)
}