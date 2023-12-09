func areNumbersAscending(s string) bool {
    words := strings.Split(s, " ")
    prev := -1

    for _, word := range words {
        curr, err := strconv.Atoi(word)
        if err != nil {
            continue
        }
        if curr <= prev {
            return false
        }
        prev = curr
    }
    return true
}