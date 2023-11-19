func vowelStrings(words []string, left int, right int) int {
    result := 0
    for i := left; i <= right; i++ {
        sChar := words[i][0]
        eChar := words[i][len(words[i]) - 1]
        if (sChar == 'a' || sChar == 'e' || sChar == 'o' || sChar == 'i' || sChar == 'u') && 
            (eChar == 'a' || eChar == 'e' || eChar == 'o' || eChar == 'i' || eChar == 'u') {
                result += 1
            }
    }
    return result
}