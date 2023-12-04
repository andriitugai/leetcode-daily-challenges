func numOfStrings(patterns []string, word string) int {
    result := 0
    for _, pat := range patterns {
        if strings.Contains(word, pat) {
            result += 1
        }
    }
    return result
}