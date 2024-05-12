func findPermutationDifference(s string, t string) int {
    ms, mt := make(map[byte]int), make(map[byte]int)
    for i := 0; i < len(s); i ++ {
        ms[s[i]] = i
        mt[t[i]] = i
    }
    result := 0
    for char, pos := range ms {
        result += abs(pos - mt[char])
    }
    return result
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}