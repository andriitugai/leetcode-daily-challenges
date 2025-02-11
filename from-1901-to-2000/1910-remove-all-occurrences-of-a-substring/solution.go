func removeOccurrences(s string, part string) string {
    for {
        idx := strings.Index(s, part)
        if idx == -1 {
            break
        }
        s = s[:idx] + s[idx + len(part):]
    }
    return s
}