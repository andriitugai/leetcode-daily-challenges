func maximumLengthSubstring(s string) int {
    var isAtMostTwo func(str string) bool
    isAtMostTwo = func(str string) bool {
        cnts := make(map[byte]int)
        var b byte
        for i := 0; i < len(str); i ++ {
            b = str[i]
            cnts[b] += 1
            if cnts[b] == 3 {
                return false
            }
        }
        return true
    }
    maxLen := 0
    for i := 0; i < len(s); i ++ {
        for j := i; j < len(s); j ++ {
            if isAtMostTwo(s[i:j + 1]) && j - i + 1> maxLen {
                maxLen = j - i + 1
            }
        }
    }
    return maxLen
}