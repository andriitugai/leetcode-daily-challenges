func beautifulSubstrings(s string, k int) int {
    v, c := make([]int, len(s) + 1), make([]int, len(s) + 1)
    for i := 0; i < len(s); i ++ {
        if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
            v[i + 1] = v[i] + 1
            c[i + 1] = c[i]
        } else {
            v[i + 1] = v[i]
            c[i + 1] = c[i] + 1
        }
    }
    result := 0
    for i := 1; i < len(s) + 1; i ++ {
        for j := 0; j < i; j ++ {
            nc := c[i] - c[j]
            nv := v[i] - v[j]
            if nc == nv && nc * nv % k == 0 {
                result += 1
            }
        }
    }
    return result
}