func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindrome(word) {
            return word
        }
    }
    return ""
}

func isPalindrome(word string) bool {
    left, right  := 0, len(word) - 1
    for left < right {
        if word[left] != word[right] {
            return false
        }
        left += 1
        right -= 1
    }
    return true
}