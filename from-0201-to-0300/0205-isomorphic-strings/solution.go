func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mapS, mapT := make(map[byte]byte), make(map[byte]byte)


    for i := 0; i < len(s); i ++ {
        if cs, ok := mapS[s[i]]; ok {
            if cs != t[i] {
                return false
            }
        } else {
            mapS[s[i]] = t[i]
        }

        if ct, ok := mapT[t[i]]; ok {
            if ct != s[i] {
                return false
            }
        } else {
            mapT[t[i]] = s[i]
        }
    }

    return true
}