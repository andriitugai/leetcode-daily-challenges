func stringHash(s string, k int) string {
    hashedChar := func(ss string) byte {
        sum := 0
        for i := 0; i < len(ss); i ++ {
            sum += int(ss[i] - 'a')
        }
        return byte(sum % 26) + 'a'
    }

    result := make([]byte, len(s) / k)
    for i := 0; i < len(s) / k; i ++ {
        result[i] = hashedChar(s[i * k : (i + 1) * k])
    }
    return string(result)
}