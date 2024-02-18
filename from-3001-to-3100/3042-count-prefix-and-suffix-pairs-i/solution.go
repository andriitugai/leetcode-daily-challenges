func countPrefixSuffixPairs(words []string) int {
    isPrefixAndSuffix := func(s1, s2 string)bool {
        n1, n2 := len(s1), len(s2)
        if n1 <= n2 && s1 == s2[:n1] && s1 == s2[n2-n1:] {
            return true
        }
        return false
    }
    result := 0
    for i := 0; i < len(words) - 1; i ++ {
        for j := i + 1; j < len(words); j ++ {
            if isPrefixAndSuffix(words[i], words[j]) {
                result += 1
            }
        }
    }
    return result
}