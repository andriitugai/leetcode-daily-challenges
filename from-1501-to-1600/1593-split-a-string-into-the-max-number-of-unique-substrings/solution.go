func maxUniqueSplit(s string) int {
    var dfs func(idx int, currSet map[string]bool) int
    dfs = func(idx int, currSet map[string]bool) int {
        if idx == len(s) {
            return 0
        }
        result := 0
        for j := idx; j < len(s); j ++ {
            substr := s[idx: j + 1]
            if currSet[substr] {
                continue
            }
            currSet[substr] = true
            result = max(result, dfs(j + 1, currSet) + 1)
            currSet[substr] = false
        }
        return result
    }
    return dfs(0, make(map[string]bool))
}