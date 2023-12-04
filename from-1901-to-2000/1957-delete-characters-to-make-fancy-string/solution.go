func makeFancyString(s string) string {
    r := []byte{}
    for i := 0; i < len(s); i ++ {
        if i > 1 && s[i] == s[i-1] && s[i] == s[i-2] {
            continue
        }
        r = append(r, s[i])
    }
    return string(r)
}