func doesAliceWin(s string) bool {
    numOfVowels := 0
    for i := 0; i < len(s); i ++ {
        if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
            numOfVowels += 1
        }
    }
    return numOfVowels > 0
}