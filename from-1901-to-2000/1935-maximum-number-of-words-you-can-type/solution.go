import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
    words := strings.Split(text, " ")
    result := 0

    for _, word := range words {
        isOk := true
        for _, letter := range brokenLetters {
            if strings.Contains(word, string(letter)) {
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