func clearDigits(s string) string {
    marks := make([]int, len(s))
    for i := 0; i < len(s); i ++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            marks[i] = 1
        } else {
            for j := i - 1; j >= 0; j -- {
                if marks[j] == 1 {
                    marks[j] = 0
                    break
                }
            }
        }
    }

    res := []string{}
    for i := 0; i < len(s); i ++ {
        if marks[i] == 1 {
            res = append(res, string(s[i]))
        }
    }
    return strings.Join(res, "")
}