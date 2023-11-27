func maximumNumberOfStringPairs(words []string) int {
    revWord := func(s string) string {
        n := len(s)
        r := make([]byte, n)
        for i := 0; i < n; i++ {
            r[i] = s[n - i - 1]
        }
        return string(r)
    }
    result := 0
    for i := 0; i < len(words) - 1; i ++ {
        reversed := revWord(words[i])
        for j := i + 1; j < len(words); j ++ {
            if reversed == words[j] {
                result += 1
            }
        }
    }
    return result
}