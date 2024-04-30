func wonderfulSubstrings(word string) int64 {
    var result int64 = 0
    count := make(map[int]int64)
    count[0] = 1
    var mask int = 0
    for i := 0; i < len(word); i ++ {
        mask ^= 1 << int(word[i] - 'a')
        result += count[mask]
        for j := 0; j < 10; j ++ {
            result += count[mask ^ (1 << j)]
        }
        count[mask] += 1
    }
    return result
}