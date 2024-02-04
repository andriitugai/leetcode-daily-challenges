func minimumTimeToInitialState(word string, k int) int {
    result := 1
    start, end := k, len(word) - k 
    for word[start:] != word[:end] {
        start += k
        if start > len(word) {
            start = len(word)
        }
        end -= k
        if end < 0 {
            end = 0
        }
        result += 1
    }
    return result
}