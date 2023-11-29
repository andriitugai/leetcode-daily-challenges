func findWordsContaining(words []string, x byte) []int {
    result := []int{}
    for i := 0; i < len(words); i ++ {
        for j := 0; j < len(words[i]); j ++ {
            if words[i][j] == x {
                result = append(result, i)
                break
            }
        }
    }
    return result
}