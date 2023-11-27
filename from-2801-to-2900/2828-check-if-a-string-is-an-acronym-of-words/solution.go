func isAcronym(words []string, s string) bool {
    acronym := make([]byte, len(words))
    for i := 0; i < len(words); i ++ {
        acronym[i] = words[i][0]
    }
    return string(acronym) == s
}