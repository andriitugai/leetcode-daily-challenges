func minRemoveToMakeValid(s string) string {
    open := 0
    sl := []byte(s)
    for i := 0; i < len(sl); i ++ {
        if sl[i] == '(' {
            open += 1
        } else if sl[i] == ')' {
            if open == 0 {
                sl[i] = '*'
            } else {
                open -= 1
            }
        }
    }
    for i := len(sl) - 1; i >= 0 && open > 0; i -- {
        if sl[i] == '(' && open > 0 {
            sl[i] = '*'
            open -=1
        }
    }
    result := []string{}
    for i := 0; i < len(sl); i ++ {
        if sl[i] != '*' {
            result = append(result, string(sl[i]))
        }
    }
    return strings.Join(result, "")
}