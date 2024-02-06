func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _, word := range strs {
        key := makeKey(word)
        groups[key] = append(groups[key], word)
    }
    result := [][]string{}
    for _, g := range(groups){
        result = append(result, g)
    }
    return result
}

func makeKey(s string) string {
    kstr := make([]byte, len(s))
    for i := 0; i < len(s); i ++ {
        kstr = append(kstr, s[i])
    } 
    sort.Slice(kstr, func(i, j int) bool{
        return kstr[i] < kstr[j]
    })
    return string(kstr)
}