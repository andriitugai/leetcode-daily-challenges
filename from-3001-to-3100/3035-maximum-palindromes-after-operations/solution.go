func maxPalindromesAfterOperations(words []string) int {
    counts := make(map[byte]int)
    sizes := make([]int, len(words))
    for k, word := range words {
        for i := 0; i < len(word); i ++ {
            counts[word[i]] += 1
        }
        sizes[k] = len(word)
    }
    pairs := 0
    for _, v := range counts {
        pairs += v / 2
    }
    sort.Ints(sizes)

    result := 0
    for _, size := range sizes {
        pairs -= size / 2
        if pairs < 0 {
            break
        }
        result += 1
    }

    return result
}