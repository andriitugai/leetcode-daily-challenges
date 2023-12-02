func reversePrefix(word string, ch byte) string {
    for i := 0; i < len(word); i ++ {
        if word[i] == ch {
            result := []byte{}
            for j := i; j >= 0; j -- {
                result = append(result, word[j])
            }
            for j := i + 1; j < len(word); j ++ {
                result = append(result, word[j])
            }
            return string(result)
        }
    }
    return word
}