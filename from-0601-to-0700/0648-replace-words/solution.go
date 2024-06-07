func replaceWords(dictionary []string, sentence string) string {
    words := strings.Split(sentence, " ")
    sort.Slice(dictionary, func(i, j int) bool {
        return len(dictionary[i]) < len(dictionary[j])
    })
    for i, word := range words {
        for _, root := range dictionary {
            if strings.HasPrefix(word, root) {
                words[i] = root
                break
            }
        }
    }
    return strings.Join(words, " ")
}