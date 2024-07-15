func getSmallestString(s string) string {
    res := []byte(s)
    for i := 1; i < len(s); i++ {
        a, b := s[i - 1], s[i]
        if a % 2 == b % 2 && a > b {
            res[i - 1], res[i] = res[i], res[i - 1]
            return string(res)
        }
    }
    return s
}