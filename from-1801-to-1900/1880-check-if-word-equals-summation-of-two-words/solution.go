func numVal(s string) int {
    n := 0
    for i := 0; i < len(s); i ++ {
        n = n * 10 + int(s[i] - 'a')
    }
    return n
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    return numVal(firstWord) + numVal(secondWord) == numVal(targetWord)
}