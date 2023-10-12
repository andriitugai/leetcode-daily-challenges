func countConsistentStrings(allowed string, words []string) int {
    allowedMap := make(map[rune]bool)
    for _, r := range allowed {
        allowedMap[r] = true
    }

    result := 0
    for _, word := range words {
        isOk := true
        for _, r := range word {
            if !allowedMap[r] {
                isOk = false
                break
            }
        }
        if isOk {
            result += 1
        }
    }

    return result
}