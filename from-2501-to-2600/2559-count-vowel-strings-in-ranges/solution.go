func vowelStrings(words []string, queries [][]int) []int {
    prefix := make([]int, len(words) + 1)
    for i, word := range words {
        prefix[i + 1] = prefix[i]
        if isVowel(word[0]) && isVowel(word[len(word) - 1]){
            prefix[i + 1] += 1
        }
    }

    result := make([]int, len(queries))
    for i, query := range queries {
        left, right := query[0], query[1]
        result[i] = prefix[right + 1] - prefix[left]
    }
    return result
}

func isVowel(c byte) bool {
    if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
        return true
    }
    return false
}