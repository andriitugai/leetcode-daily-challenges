func addMinimum(word string) int {
    pattern := "abc"
    pp := 0
    idx := 0
    result := 0
    for idx < len(word) {
        if pp == len(pattern) {
            pp = 0
        }
        if word[idx] == pattern[pp] {
            idx += 1
        } else {
            result += 1
        }
        pp += 1
    }
    return result - pp + len(pattern)
}