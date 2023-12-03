func numDifferentIntegers(word string) int {
    numbers := make(map[string]bool)
    si := -1
    for i := 0; i < len(word); i ++ {
        c := word[i]
        if c >= '0' && c <= '9' {
            if si == -1 {
                si = i
            }
        } else {
            if si != -1 {
                for word[si] == '0' && si < i - 1 {
                    si += 1
                }
                numbers[word[si:i]] = true
                si = -1
            }
        }
    }
    if si != -1 {
        for word[si] == '0' && si < len(word) - 1 {
            si += 1
        }
        numbers[word[si:]] = true
    }
    return len(numbers)
}