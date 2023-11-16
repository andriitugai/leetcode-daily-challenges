func similarPairs(words []string) int {
    wordChars := make(map[string][]bool)
    for _, word := range words {
        wordChars[word] = make([]bool, 26)
        for _, char := range word {
            wordChars[word][char - 'a'] = true
        }
    }

    result := 0
    for i := 0; i < len(words); i ++ {
        for j := i + 1; j < len(words); j ++ {
            if isSimilar(wordChars[words[i]], wordChars[words[j]]) {
                result += 1
            }
        }
    }
    return result
}

func isSimilar(c1, c2 []bool) bool {
    for i := 0; i < 26; i ++ {
        if c1[i] != c2[i] {
            return false
        }
    }
    return true
}