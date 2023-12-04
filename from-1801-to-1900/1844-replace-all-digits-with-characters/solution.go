func replaceDigits(s string) string {
    result := []byte{}
    for i := 0; i < len(s); i += 2 {
        result = append(result, s[i])
        if i < len(s) - 1 {
            result = append(result, s[i] + s[i+1] - '0')
        }   
    }
    return string(result)
}