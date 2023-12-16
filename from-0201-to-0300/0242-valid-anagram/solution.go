func isAnagram(s string, t string) bool {
    ct, cs := make(map[rune]int), make(map[rune]int)

    for _, c := range s {
        cs[c] += 1
    }
    for _, c := range t {
        ct[c] += 1
    }

    for c, v := range cs {
        if ct[c] != v {
            return false
        }
    }

    for c, v := range ct {
        if cs[c] != v {
            return false
        }
    }
    return true
}