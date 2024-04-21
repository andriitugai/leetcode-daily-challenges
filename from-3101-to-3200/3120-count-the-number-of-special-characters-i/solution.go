func numberOfSpecialChars(word string) int {
    diff := byte('a' - 'A')
    seen := make(map[byte]bool)
    result := 0

    for i := 0; i < len(word); i ++ {
        seen[word[i]] = true
    }

    for k, _ := range seen {
        if k > 'Z' && seen[k - diff] {
            result += 1
        }
    }
    return result
}