func findTheLongestBalancedSubstring(s string) int {
    maxBalanced := 0

    for i := 0; i < len(s); {
        curZeros := 0
        for i < len(s) && s[i] == '0' {
            curZeros += 1
            i += 1
        }
        curOnes := 0
        for i < len(s) && s[i] == '1' {
            curOnes += 1
            i += 1
        }
        if curZeros != 0 {
            row := curZeros
            if curOnes < row {
                row = curOnes
            }
            if maxBalanced < 2 * row {
                maxBalanced = 2 * row
            }
        }
    }
    return maxBalanced
}