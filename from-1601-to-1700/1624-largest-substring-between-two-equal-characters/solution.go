func maxLengthBetweenEqualCharacters(s string) int {
    result := -1
    pos := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        c := s[i]
        if k, exists := pos[c]; !exists {
            pos[c] = i
        } else {
            dist := i - k - 1
            if dist > result {
                result = dist
            }
        }
    }
    return result
}