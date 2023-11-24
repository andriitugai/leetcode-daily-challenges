func splitWordsBySeparator(words []string, separator byte) []string {
    result := []string{}
    for _, word := range words {
        for _, w := range strings.Split(word, string(separator)) {
            if w != "" {
                result = append(result, w)
            }
        }
    }
    return result
}