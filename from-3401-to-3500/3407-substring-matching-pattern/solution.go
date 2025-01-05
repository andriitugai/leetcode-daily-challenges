func hasMatch(s string, p string) bool {
    idx := strings.Index(p, "*")
    p1, p2 := p[:idx], p[idx+1:]

    i1 := strings.Index(s, p1)
    if i1 < 0 {
        return false
    }
    return strings.Index(s[i1 + len(p1):], p2) >= 0
}