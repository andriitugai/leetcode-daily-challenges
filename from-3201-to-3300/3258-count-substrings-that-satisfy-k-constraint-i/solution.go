func countKConstraintSubstrings(s string, k int) int {
    n := len(s)
    result := 0
    for i := 0; i < n; i ++ {
        cnt1, cnt0 := 0, 0
        for j := i; j < n; j ++ {
            if s[j] == '0' {
                cnt0 += 1
            } else {
                cnt1 += 1
            }
            if cnt1 <= k || cnt0 <= k {
                result += 1
            } else {
                break
            }
        }
    }
    return result
}