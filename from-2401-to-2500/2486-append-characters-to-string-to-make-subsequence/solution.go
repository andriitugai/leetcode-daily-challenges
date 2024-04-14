func appendCharacters(s string, t string) int {
    p := 0
    for i := 0; i < len(s) && p < len(t); i ++ {
        if s[i] == t[p] {
            p += 1
        }
    }
    return len(t) - p
}