func isSubstringPresent(s string) bool {
    var b1, b2 byte
    n := len(s)
    for i := 0; i < n - 1; i ++ {
        b2, b1 = s[i], s[i+1]
        for j := i; j < n - 1; j ++ {
            if s[j] == b1 && s[j+1] == b2 {
                return true
            }
        }
    }
    return false
}