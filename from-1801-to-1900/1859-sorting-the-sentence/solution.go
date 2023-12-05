func sortSentence(s string) string {
    words := strings.Split(s, " ")
    sort.Slice(words, func(i, j int) bool {
        return words[i][len(words[i])-1] < words[j][len(words[j])-1]
    })
    result := words[0][:len(words[0])-1]
    for i := 1; i < len(words); i ++ {
        result += " " + words[i][:len(words[i])-1]
    }
    return result
}