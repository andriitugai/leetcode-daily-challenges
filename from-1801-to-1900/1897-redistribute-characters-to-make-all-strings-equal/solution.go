func makeEqual(words []string) bool {
    chars := make(map[byte]int)
    for _, word := range words {
        for i := 0; i < len(word); i ++ {
            chars[word[i]] += 1
        }
    }
    for _, cnt := range chars {
        if cnt % len(words) != 0 {
            return false
        }
    }
    return true
}