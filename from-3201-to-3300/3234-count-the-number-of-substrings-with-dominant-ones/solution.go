func numberOfSubstrings(s string) int {
    n := len(s)
    result := 0
    zeros := []int{}
    for i := 0; i < n; i ++ {
        if s[i] == '0' {
            zeros = append(zeros, i)
        }

        l := 0
        for l <= len(zeros) {
            targetLen := l * l + l
            if targetLen > i + 1 {
                break
            }

            leftIdx := -1
            if l < len(zeros) {
                leftIdx = zeros[len(zeros) - 1 - l]
            }
            rightIdx := i
            if l > 0 {
                rightIdx = zeros[len(zeros) - l]
            }

            minLen := i - rightIdx + 1
            maxLen := i - leftIdx
            if minLen >= targetLen {
                result += maxLen - minLen + 1
            } else if minLen <= targetLen && targetLen <= maxLen {
                result += maxLen - targetLen + 1
            }

            l += 1
        }
    }
    return result
}