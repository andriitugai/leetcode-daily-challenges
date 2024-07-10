func validStrings(n int) []string {
    result := []string{}
    var helper func(i int, curr []byte)
    helper = func(i int, curr []byte) {
        if i == n {
            result = append(result, string(curr))
            return
        }
        if curr[i - 1] == '0' {
            helper(i + 1, append(curr, '1'))
        }
        if curr[i - 1] == '1' {
            helper(i + 1, append(curr, '0'))
            helper(i + 1, append(curr, '1'))
        }
    }
    helper(1, []byte{'0'})
    helper(1, []byte{'1'})
    return result
}