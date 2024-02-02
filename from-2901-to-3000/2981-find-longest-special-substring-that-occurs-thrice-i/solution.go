func maximumLength(s string) int {
    maxLen := -1
    simSet := make(map[byte]bool)
    for i := 0; i < len(s); i ++ {
        simSet[s[i]] = true
    }
    abc := []byte{}
    for sim, _ := range simSet {
        abc = append(abc, sim)
    }

    makeSpecialString := func(b byte, n int) string {
        bytes := make([]byte, n)
        for i := 0; i < n; i ++ {
            bytes[i] = b
        }
        return string(bytes)
    }
    for sLen := 1; sLen < len(s); sLen ++ {
        for _, sim := range abc {
            ss := makeSpecialString(sim, sLen)
            count := 0
            for start := 0; start < len(s) - sLen + 1; start ++ {
                if s[start:start+sLen] == ss {
                    count += 1
                    if count == 3 {
                        maxLen = sLen
                        break
                    }
                }
            }
            if maxLen == sLen {
                break
            }
        }
    }
    return maxLen
}