func maxOperations(s string) int {
    n := len(s)
    result := 0
    ones := 0
    for i := 0; i < n; i ++ {
        if s[i] == '1' {
            ones += 1
        }
        if i > 0 && s[i] < s[i - 1] {
            result += ones
        }
    }
    return result
}