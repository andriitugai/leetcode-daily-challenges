func getSmallestString(s string, k int) string {
    ss := []byte(s)

    charDistance := func(c byte) int {
        left := int(c - 'a')
        right := int('z' - c + 1)
        if left < right {
            return left
        }
        return right
    }

    for i, c := range ss {
        dist := charDistance(c)
        if k >= dist {
            k -= dist
            ss[i] = 'a'
        } else {
            ss[i] = ss[i] - byte(k)
            k = 0
        }
    }

    return string(ss)
}