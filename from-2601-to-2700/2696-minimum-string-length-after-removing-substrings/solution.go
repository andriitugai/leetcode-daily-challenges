func minLength(s string) int {
    for true {
        curlen := len(s)
        s = strings.ReplaceAll(s, "AB", "")
        s = strings.ReplaceAll(s, "CD", "")
        if len(s) == curlen {
            break
        }
    }
    return len(s)
}