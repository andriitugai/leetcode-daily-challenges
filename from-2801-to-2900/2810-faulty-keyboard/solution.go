func finalString(s string) string {
    result := []byte{}
    for i := 0; i < len(s); i++ {
        if s[i] != 'i'{
            result = append(result, s[i])
        } else {
            l, r := 0, len(result) - 1
            for l < r {
                result[l], result[r] = result[r], result[l]
                l += 1
                r -= 1
            }
        }
    }
    return string(result)
}