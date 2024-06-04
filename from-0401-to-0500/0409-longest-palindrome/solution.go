func longestPalindrome(s string) int {
    cnt := make(map[byte]int)
    for i := 0; i < len(s); i ++ {
        cnt[s[i]] += 1
    }

    result := 0
    maxOdd := 0
    for _, val := range cnt {
        if val % 2 == 0 {
            result += val
        } else {
            result += val - 1
            if val > maxOdd {
                maxOdd = 1
            }
        }
    }
    return result + maxOdd
}