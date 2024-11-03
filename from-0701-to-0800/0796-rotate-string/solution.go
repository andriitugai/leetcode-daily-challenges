func rotateString(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }
    for k := 0; k < len(s); k ++ {
        if s[k:] + s[:k] == goal {
            return true
        }
    }
    return false
}