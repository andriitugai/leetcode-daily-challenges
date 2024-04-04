func maxDepth(s string) int {
    result, deep := 0, 0
    for i := 0; i < len(s); i ++ {
        if s[i] == '(' {
            deep += 1
            if deep > result {
                result = deep
            } 
        } else if s[i] == ')' {
            deep -= 1
        }
    }
    return result
}