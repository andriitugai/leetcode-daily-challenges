func maximumOddBinaryNumber(s string) string {
    n := len(s)
    ones := 0
    for i := 0; i < n; i ++ {
        if s[i] == '1' {
            ones += 1
        }
    }
    r := make([]byte, n)
    i := 0
    for i < ones - 1 {
        r[i] = '1'
        i += 1
    }
    for i < n - 1 {
        r[i] = '0'
        i += 1
    }
    r[n-1] = '1'
    return string(r)
}