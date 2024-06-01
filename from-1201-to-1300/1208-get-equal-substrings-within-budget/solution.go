func equalSubstring(s string, t string, maxCost int) int {
    result := 0
    cost := 0
    left := 0

    for right := 0; right < len(s); right ++ {
        cost += costToChange(s[right], t[right])
        for cost > maxCost {
            cost -= costToChange(s[left], t[left])
            left += 1
        }
        result = max(result, right - left + 1)
    }
    return result
}

func costToChange(a, b byte) int {
    if b > a {
        a, b = b, a
    }
    return int(a) - int(b)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}