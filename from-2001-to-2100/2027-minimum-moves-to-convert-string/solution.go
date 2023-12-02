func minimumMoves(s string) int {
    result := 0
    i := 0
    for i < len(s) {
        if s[i] == 'X' {
            result += 1
            i += 3
        } else {
            i += 1
        }
    }
    return result
}