func scoreOfString(s string) int {
    result := 0
    for i := 1; i < len(s); i ++ {
        result += absDiff(s[i - 1], s[i])
    }
    return result
}

func absDiff(a, b byte) int {
    if a < b {
        return int(b - a)
    }
    return int(a - b)
}