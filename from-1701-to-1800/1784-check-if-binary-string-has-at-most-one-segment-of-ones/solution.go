func checkOnesSegment(s string) bool {
    zeros := false
    for i := 0; i < len(s); i ++ {
        if s[i] == '0' {
            zeros = true
        } else if zeros && s[i] == '1' {
            return false
        }
    }
    return true
}