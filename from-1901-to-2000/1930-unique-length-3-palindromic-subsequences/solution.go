func countPalindromicSubsequence(s string) int {
    // Define the first and the last index for every letter in the alphabet
    firstIndex := make([]int, 26)
    lastIndex := make([]int, 26)
    for i := 0; i < 26; i ++ {
        firstIndex[i] = -1
        lastIndex[i] = -1
    }

    for i := 0; i < len(s); i ++ {
        idx := int(s[i]-'a')
        if firstIndex[idx] == -1 {
            firstIndex[idx] = i
            lastIndex[idx] = i
        } else {
            lastIndex[idx] = i
        }
    }

    result := 0
    // for every letter in the alphabet
    for i := 0; i < 26; i ++ {
        fi, li := firstIndex[i], lastIndex[i]
        // if the letter occurs more than once and these occurencies are not next to each other
        if fi >= 0 && li - fi > 1 {
            // check every other letter
            for j := 0; j < 26; j ++ {
                // if it occurs between the first and last indexes
                for k := fi + 1; k < li; k ++ {
                    if int(s[k] - 'a') == j {
                        result += 1
                        break
                    }
                }
            }
        }
    }

    return result
}