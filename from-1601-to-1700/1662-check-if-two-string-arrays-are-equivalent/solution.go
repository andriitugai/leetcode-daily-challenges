unc arrayStringsAreEqual(word1 []string, word2 []string) bool {
    i1, i2 := 0, 0
    j1, j2 := 0, 0
    m1, m2 := len(word1), len(word2)
    n1, n2 := len(word1[0]), len(word2[0])

    for i1 < m1 && i2 < m2 {
        if word1[i1][j1] != word2[i2][j2] {
            return false
        }
        j1 += 1
        if j1 == n1 {
            i1 += 1
            j1 = 0
            if i1 < m1 {
                n1 = len(word1[i1])
            }
        }
        j2 += 1
        if j2 == n2 {
            i2 += 1
            j2 = 0
            if i2 < m2 {
                n2 = len(word2[i2])
            }
        }
    }
    return i1 == m1 && i2 == m2
}