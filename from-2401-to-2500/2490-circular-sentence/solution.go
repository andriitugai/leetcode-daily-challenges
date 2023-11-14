func isCircularSentence(sentence string) bool {
    if sentence[0] != sentence[len(sentence)-1] {
        return false
    }
    words := strings.Fields(sentence)
    for i := 0; i < len(words) - 1; i++ {
        if words[i][len(words[i])-1] != words[i+1][0] {
            return false
        }
    }
    return true
}