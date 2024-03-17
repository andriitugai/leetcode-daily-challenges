func minimizeStringValue(s string) string {
    n := len(s)
    charFreq := make([]int, 26)
    temp := []string{}
    for i := 0; i < n; i ++ {
        if s[i] != '?' {
            charFreq[s[i] - 'a'] += 1
        }
    }
    for i := 0; i < n; i ++ {
        if s[i] != '?' {
            continue
        }
        minFreqIdx, minFreq := 0, 1000001
        for j := 0; j < 26; j ++ {
            if minFreq > charFreq[j] {
                minFreq, minFreqIdx = charFreq[j], j
            }
        }
        temp = append(temp, string('a' + minFreqIdx))
        charFreq[minFreqIdx] += 1
    }
    sort.Strings(temp)
    result := make([]string, n)
    tmpIdx := 0
    for i := 0; i < n; i ++ {
        if s[i] == '?' {
            result[i] = temp[tmpIdx]
            tmpIdx += 1
        } else {
            result[i] = string(s[i])
        }
    }
    return strings.Join(result, "")
}