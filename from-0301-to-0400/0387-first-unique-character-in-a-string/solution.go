func firstUniqChar(s string) int {
    pos := make([]int, 26)
    for i := 0; i < len(s); i ++ {
        if pos[int(s[i] - 'a')] == 0 {
            pos[int(s[i] - 'a')] = i + 1
        } else {
            pos[int(s[i] - 'a')] = -1
        }
    }
    minIdx := 1000001
    for i := 0; i < 26; i ++ {
        if pos[i] >0 && pos[i] < minIdx {
            minIdx = pos[i]
        }
    }
    if minIdx == 1000001 {
        return -1
    }
    return minIdx - 1
}