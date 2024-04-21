func numberOfSpecialChars(word string) int {
    result := 0
    lPos, uPos := make([]int, 26), make([]int, 26)
    for i := 0; i < 26; i ++ {
        lPos[i] = -1
        uPos[i] = -1
    }

    for i := 0; i < len(word); i ++ {
        if word[i] > 'Z' {
            lPos[word[i] - 'a'] = i
        } else if uPos[word[i] - 'A'] == -1 {
            uPos[word[i] - 'A'] = i
        }
    }

    for i := 0; i < 26; i ++ {
        if lPos[i] >= 0 && lPos[i] < uPos[i] {
            result += 1
        }
    }
    
    return result
}