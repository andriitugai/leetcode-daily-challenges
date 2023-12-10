func divideString(s string, k int, fill byte) []string {
    if len(s) % k != 0 {
        s = s + strings.Repeat(string(fill), k - len(s) % k)
    }
    result := []string{}
    for i := 0; i < len(s); i += k {
        result = append(result, s[i:i+k])
    }
    return result
}