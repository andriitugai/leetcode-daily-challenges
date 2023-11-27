func canBeEqual(s1 string, s2 string) bool {
    return string([]byte{s1[2], s1[1], s1[0], s1[3]}) == s2 ||
            string([]byte{s1[0], s1[3], s1[2], s1[1]}) == s2 ||
            string([]byte{s1[2], s1[3], s1[0], s1[1]}) == s2 ||
            s1 == s2
}