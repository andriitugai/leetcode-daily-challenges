type charPos struct {
    char byte
    pos int
}

func lastNonEmptyString(s string) string {
    maxCount := 0
    counts := make(map[byte]int)
    for i := 0; i < len(s); i ++ {
        counts[s[i]] += 1
    }

    for _, cnt := range counts {
        if cnt > maxCount {
            maxCount = cnt
        }
    }

    res := []charPos{}
    for b, cnt := range counts {
        if cnt == maxCount {
            res = append(res, charPos{char: b, pos: strings.LastIndexByte(s, b)})
        }
    }

    sort.Slice(res, func(i, j int) bool {
        return res[i].pos < res[j].pos
    })

    ans := []byte{}
    for _, cp := range res {
        ans = append(ans, cp.char)
    }
    return string(ans)
}