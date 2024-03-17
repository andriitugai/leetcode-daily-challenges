func countSubstrings(s string, c byte) int64 {
    var cc int64 = 0
    for i := 0; i < len(s); i ++ {
        if s[i] == c {
            cc += 1
        }
    }
    return cc * (cc + 1) / 2
}