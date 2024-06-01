func compressedString(word string) string {
    result := []string{}
    currChar := word[0]
    currCount := 1
    flush := func() {
        if currCount > 0 {
            result = append(result, fmt.Sprintf("%d%s", currCount, string(currChar)))
        }
    }
    for i := 1; i < len(word); i ++ {
        if word[i] != currChar {
            flush()
            currChar = word[i]
            currCount = 1
        } else {
            currCount += 1
            if currCount == 9 {
                flush()
                currCount = 0
            }
        }
    }
    flush()
    return strings.Join(result, "")
}