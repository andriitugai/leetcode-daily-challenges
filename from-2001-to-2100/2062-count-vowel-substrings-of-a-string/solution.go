unc countVowelSubstrings(word string) int {
    result := 0
    for i:=0; i<len(word); i++ {
        check := make(map[byte]int)
        for j:=i; j<len(word); j++ {
            if strings.Contains("aeiou", string(word[j])) {
                check[word[j]] += 1
            } else {
                break
            }
            if len(check) == 5 {
                result += 1
            }
        }
    }
    return result
}