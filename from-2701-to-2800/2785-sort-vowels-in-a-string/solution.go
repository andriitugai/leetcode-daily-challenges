func sortVowels(s string) string {
    buffer := []byte{}
    for i:=0; i<len(s); i++ {
        if isVowel(s[i]) {
            buffer = append(buffer, s[i])
        }
    }

    sort.Slice(buffer, func(i int, j int) bool { 
        return buffer[i] < buffer[j]
    })

    result := []byte(s)
    buffer_idx := 0
    for i:=0; i<len(s); i++ {
        if isVowel(s[i]) {
            result[i] = buffer[buffer_idx]
            buffer_idx += 1
        }
    }

    return string(result)
}

func isVowel(s byte) bool{
    return s=='a'|| s=='e'|| s=='u'|| s=='o'|| s=='i'|| s =='A' || s=='E'|| s=='U'||s =='O'|| s=='I'
}