func commonChars(words []string) []string {
    counters := make([]map[byte]int, len(words))
    var cnt map[byte]int
    for i, word := range words {
        cnt = make(map[byte]int)
        for i := 0; i < len(word); i ++ {
            cnt[word[i]] += 1
        }
        counters[i] = cnt
    }

    result := []string{}
    for char, numChars := range counters[0] {
        minChars := numChars
        for _, cnt := range counters[1:] {
            num := cnt[char]
            if num < minChars {
                minChars = num
            }
            if minChars == 0 {
                break
            }
        }
        for i := 0; i < minChars; i ++ {
            result = append(result, string(char))
        }
    }
    return result
}