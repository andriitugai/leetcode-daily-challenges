func punishmentNumber(n int) int {
    var canSplit func(s string, tgt int) bool 
    canSplit = func(s string, tgt int) bool {
        if tgt < 0 {
            return false
        }
        if len(s) == 0 && tgt == 0 {
            return true
        }
        ans := false
        for i := 0; i < len(s); i ++ {
            leftNum, _ := strconv.Atoi(s[:i + 1])
            rightStr := s[i + 1:]
            if canSplit(rightStr, tgt - leftNum) {
                ans = true
                return ans
            }
        }
        return ans
    }
    var square int
    result := 0
    for i := 1; i <= n; i ++ {
        square = i * i
        if canSplit(strconv.Itoa(square), i){
            result += square
        }
    }
    return result
}