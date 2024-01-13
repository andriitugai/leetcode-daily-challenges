func minSteps(s string, t string) int {
    cs := make(map[byte]int)
    ct := make(map[byte]int)
    for i := 0; i < len(s); i ++ {
        cs[s[i]] += 1
        ct[t[i]] += 1
    }

    result := 0
    for k, v := range ct {
        if v > cs[k]{
            result += (v - cs[k])
        }
    }
    return result
}