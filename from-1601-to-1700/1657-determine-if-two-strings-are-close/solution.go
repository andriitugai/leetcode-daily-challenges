unc closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }
    m1 := make([]int, 26)
    m2 := make([]int, 26)
    for i := 0; i < len(word1); i ++ {
        m1[int(word1[i] - 'a')] += 1
    }
    for i := 0; i < len(word2); i ++ {
        m2[int(word2[i] - 'a')] += 1
    }
    for i := 0; i < 26; i ++ {
        if m1[i] > 0 && m2[i] == 0 || m1[i] == 0 && m2[i] > 0 {
            return false
        }
    }
    sort.Ints(m1)
    sort.Ints(m2)
    for i := 0; i < 26; i ++ {
        if m1[i] != m2[i] {
            return false
        }
    }
    return true
}