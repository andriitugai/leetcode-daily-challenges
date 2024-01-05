func removeStars(s string) string {
    n := len(s)

    stack := []byte{}
    for i := 0; i < n; i ++ {
        if s[i] == '*' {
            stack = stack[:len(stack) - 1]
        } else {
            stack = append(stack, s[i])
        }
    }

    return string(stack)
}