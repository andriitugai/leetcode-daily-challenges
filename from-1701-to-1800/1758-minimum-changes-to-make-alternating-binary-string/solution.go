func minOperations(s string) int {
    c1, c0 := 0, 0
    for i := 0; i < len(s); i ++ {
        if i % 2 == 0 {
            if s[i] == '1' {
                c0 += 1
            } else {
                c1 += 1
            }
        } else {
            if s[i] == '1' {
                c1 += 1
            } else {
                c0 += 1
            }
        }
    }
    if c1 < c0 {
        return c1
    }
    return c0
}