func maxLength(arr []string) int {
    currSet := make(map[rune]bool)
    canAdd := func(currSet map[rune]bool, candidate string) bool {
        candSet := make(map[rune]bool)
        for _, r := range candidate {
            if candSet[r] || currSet[r] {
                return false
            }
            candSet[r] = true
        }
        return true
    }
    var backtrack func(i int) int
    backtrack = func(i int) int {
        if i == len(arr) {
            return len(currSet)
        }
        var ans int = 0
        if canAdd(currSet, arr[i]) {
            for _, r := range arr[i] {
                currSet[r] = true
            }
            ans = backtrack(i + 1)
            for _, r := range arr[i] {
                delete(currSet, r)
            }
        }
        a2 := backtrack(i + 1)
        if a2 > ans {
            return a2
        }
        return ans
    }
    return backtrack(0)
}