func kthCharacter(k int) byte {
    s := "a"
    for len(s) < k {
        buff := make([]byte, len(s))
        for i := 0; i < len(s); i ++ {
            chNum := (s[i] - 'a' + 1) % 26
            buff[i] = 'a' + chNum
        }
        s += string(buff)
    }
    return s[k - 1]
}