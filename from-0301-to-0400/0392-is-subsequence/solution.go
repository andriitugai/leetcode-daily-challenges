func isSubsequence(s string, t string) bool {
    if s == "" {
        return true
    }
    if t == "" {
        return false
    }
    ps := 0
    for pt := 0; pt < len(t); pt ++ {
        if s[ps] == t[pt] {
            ps += 1
            if ps == len(s) {
                return true
            }
        }
    }
    return false
}