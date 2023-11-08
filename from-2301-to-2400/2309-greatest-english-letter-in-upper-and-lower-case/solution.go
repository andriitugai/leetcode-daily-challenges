func greatestLetter(s string) string {
    count := make([]int, 26)

    for i := 0; i < len(s); i ++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            count[int(s[i] - 'a')] |= 1
        } else {
            count[int(s[i] - 'A')] |= 2
        }
    }

    for i := 25; i >= 0; i-- {
        if count[i] == 3 {
            return string('A' + byte(i))
        }
    }
    return ""
}