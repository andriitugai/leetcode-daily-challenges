func minMaxDifference(num int) int {
    s := strconv.Itoa(num)
    tmp := s
    for _, r := range s {
        if r != rune('9') {
            tmp = strings.ReplaceAll(s, string(r), "9")
            break
        }
    }
    ans, _ := strconv.Atoi(tmp)
    tmp = s
    for _, r := range s {
        if r != rune('0') {
            tmp = strings.ReplaceAll(s, string(r), "0")
            break
        }
    }
    t, _ := strconv.Atoi(tmp)
    return ans - t
}