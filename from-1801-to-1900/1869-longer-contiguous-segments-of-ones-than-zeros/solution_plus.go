func checkZeroOnes(s string) bool {
    ones := strings.Split(s, "0")
    zeroes := strings.Split(s, "1")
    
    longestOnes, longestZeroes := 0, 0
    
    for _, part := range ones {
        if len(part) > longestOnes {
            longestOnes = len(part)
        }
    }
    
    for _, part := range zeroes {
        if len(part) > longestZeroes {
            longestZeroes = len(part)
        }
    }
    
    return longestOnes > longestZeroes
}