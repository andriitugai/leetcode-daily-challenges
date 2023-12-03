func secondHighest(s string) int {
    digits := make([]bool, 10)
    for i := 0; i < len(s); i ++ {
        if s[i] >= '0' && s[i] <= '9' {
            digits[s[i] - '0'] = true
        }
    }

    first := false
    for i := 9; i >= 0; i -- {
        if digits[i] {
            if first {
                return i
            } else {
                first = true
            }
        }
    }
    return -1
}