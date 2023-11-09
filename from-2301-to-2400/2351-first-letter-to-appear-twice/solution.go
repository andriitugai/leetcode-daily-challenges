func repeatedCharacter(s string) byte {
    cnt := make(map[byte]int)

    for i := 0; i < len(s); i++ {
        cnt[s[i]] += 1
        if cnt[s[i]] == 2 {
            return s[i]
        }
    }
    return '1'
}