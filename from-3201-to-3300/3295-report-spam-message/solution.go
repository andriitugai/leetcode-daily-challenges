func reportSpam(message []string, bannedWords []string) bool {
    banned := make(map[string]bool)
    for _, word := range bannedWords {
        banned[word] = true
    }
    banCount := 0
    for _, word := range message {
        if banned[word] {
            banCount += 1
        }
        if banCount == 2 {
            return true
        }
    }
    return false
}