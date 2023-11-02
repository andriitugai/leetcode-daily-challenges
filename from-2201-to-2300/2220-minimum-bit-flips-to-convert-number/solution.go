func minBitFlips(start int, goal int) int {
    code := start ^ goal
    count := 0
    for code > 0 {
        code = code & (code - 1)
        count += 1
    }
    return count
}