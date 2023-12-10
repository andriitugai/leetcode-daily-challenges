func checkAlmostEquivalent(word1 string, word2 string) bool {
    c1 := make(map[byte]int)
    for i := 0; i < len(word1); i ++ {
        c1[word1[i]] += 1
    }
    for i := 0; i < len(word2); i ++ {
        c1[word2[i]] -= 1
    }

    for _, diff := range c1 {
        if diff < -3 || diff > 3 {
            return false
        }
    }
    return true
}