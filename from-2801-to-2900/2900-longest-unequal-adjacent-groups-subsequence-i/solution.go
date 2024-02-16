func getLongestSubsequence(words []string, groups []int) []string {
    n := len(groups)

    getAlternative := func(startVal int) []int {
        idxs := []int{}
        i := 0
        for i < n {
            if groups[i] == startVal {
                idxs = append(idxs, i)
                break
            }
            i += 1
        }
        lastVal := startVal
        for i < n {
            if groups[i] != lastVal {
                idxs = append(idxs, i)
                lastVal = groups[i]
            }
            i += 1
        }
        return idxs
    }

    fromOne := getAlternative(1)
    fromZero := getAlternative(0)
    
    longest := fromOne
    if len(fromZero) > len(fromOne) {
        longest = fromZero
    }

    result := []string{}
    for _, idx := range longest {
        result = append(result, words[idx])
    }
    return result
}