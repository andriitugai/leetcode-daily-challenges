func maxRepeating(sequence string, word string) int {
    maxRepeat, curRepeat := 0, 0
    wLen := len(word)
    for i := 0; i <= len(sequence) - wLen; i++ {
        if sequence[i: i + wLen] == word {
            curRepeat += 1
            i += wLen - 1
        } else {
            i -= curRepeat * wLen
            curRepeat = 0
        }
        if curRepeat > maxRepeat {
            maxRepeat = curRepeat
        }
    }
    return maxRepeat
}