func shortestSubstrings(arr []string) []string {
    result := make([]string, len(arr))
    minSubPattern := "~~~~~~~~~~~~~~~~~~~~~"
    for idx, word := range arr {
        minSub := minSubPattern
        for i := 0; i < len(word); i ++ {
            for j := i; j < len(word); j ++ {
                sub := word[i:j+1]
                isUnique := true

                for k, otherWord := range arr{
                    if k == idx {
                        continue
                    }
                    if strings.Contains(otherWord, sub) {
                        isUnique = false
                        break
                    }                
                }
                if isUnique && len(sub) <= len(minSub) {
                    if sub < minSub || len(sub) < len(minSub) {
                        minSub = sub
                    }
                }
            }
        }
        if minSub == minSubPattern {
            minSub = ""
        }
        result[idx] = minSub
    }
    return result
}