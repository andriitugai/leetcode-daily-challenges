func minimumString(a string, b string, c string) string {
    minLen := 500
    var minStr string

    res := combineStrings(a, combineStrings(b, c))
    fmt.Println(res)
    if len(res) < minLen {
        minLen = len(res)
        minStr = res
    } else if len(res) == minLen && res < minStr {
        minStr = res
    }
    res = combineStrings(a, combineStrings(c, b))
    fmt.Println(res)
    if len(res) < minLen {
        minLen = len(res)
        minStr = res
    } else if len(res) == minLen && res < minStr {
        minStr = res
    }
    res = combineStrings(b, combineStrings(c, a))
    fmt.Println(res)
    if len(res) < minLen {
        minLen = len(res)
        minStr = res
    } else if len(res) == minLen && res < minStr {
        minStr = res
    }
    res = combineStrings(b, combineStrings(a, c))
    fmt.Println(res)
    if len(res) < minLen {
        minLen = len(res)
        minStr = res
    } else if len(res) == minLen && res < minStr {
        minStr = res
    }
    res = combineStrings(c, combineStrings(b, a))
    fmt.Println(res)
    if len(res) < minLen {
        minLen = len(res)
        minStr = res
    } else if len(res) == minLen && res < minStr {
        minStr = res
    }
    res = combineStrings(c, combineStrings(a, b))
    fmt.Println(res)
    if len(res) < minLen {
        minLen = len(res)
        minStr = res
    } else if len(res) == minLen && res < minStr {
        minStr = res
    }
    return minStr
}

func combineStrings(a, b string) string {
    if strings.Contains(a, b) {
        return a
    } else if strings.Contains(b, a) {
        return b
    }
    mSuf := 0
    for i := 0; i < len(b); i ++ {
        if strings.HasSuffix(a, b[:i]) {
            mSuf = i
        }
    }
    return a + b[mSuf:]
}