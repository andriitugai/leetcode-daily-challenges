func minimumSteps(s string) int64 {
    var result, zeros int64 = 0, 0
    for i := len(s) - 1; i >= 0; i -- {
        if s[i] == '0' {
            zeros += 1
        } else {
            result += zeros
        }
    }
    return result
}