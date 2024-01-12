func halvesAreAlike(s string) bool {
    n := len(s) / 2
    c1, c2 := 0, 0
    vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
    isVowel := func(c byte) bool {
        for _, v := range vowels {
            if c == v {
                return true
            }
        }
        return false
    }

    for i := 0; i < n; i ++ {
        if isVowel(s[i]) {
            c1 += 1
        }
        if isVowel(s[i + n]) {
            c2 += 1
        }
    }
    return c1 == c2
}