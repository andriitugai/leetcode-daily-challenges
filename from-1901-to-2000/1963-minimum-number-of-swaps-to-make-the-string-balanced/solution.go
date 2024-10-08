func minSwaps(s string) int {
    extra := 0
    maxExtra := 0
    for i := 0; i < len(s); i ++ {
        if s[i] == '[' {
            extra -= 1
        } else {
            extra += 1
        }
        if extra > maxExtra {
            maxExtra = extra
        }
    }
    return (maxExtra + 1) / 2
}