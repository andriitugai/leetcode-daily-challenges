func countAsterisks(s string) int {
    inside := false
    result := 0
    for i := 0; i < len(s); i ++ {
        if s[i] == '|' {
            inside = !inside
        } else if s[i] == '*' && !inside {
            result += 1
        }
    }
    return result
}