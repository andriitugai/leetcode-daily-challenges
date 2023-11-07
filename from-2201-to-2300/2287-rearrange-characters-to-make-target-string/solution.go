unc rearrangeCharacters(s string, target string) int {
    smap := make(map[byte]int)
    tmap := make(map[byte]int)

    for i := 0; i < len(s); i ++ {
        smap[s[i]] += 1
    }
    for i := 0; i < len(target); i ++ {
        tmap[target[i]] += 1
    }

    result := 1000
    for kt, vt := range tmap {
        r := smap[kt] / vt
        if r < result {
            result = r
        }
    }
    return result
}