func countHomogenous(s string) int {
    mod := 1000000007
    result := 0
    var curChar byte = 0
    curSize := 0

    for i := 0; i < len(s); i ++ {
        if s[i] != curChar {
            curChar = s[i]
            curSize = 0
        }
        curSize += 1
        result = (result + curSize) % mod
    }
    return result
}