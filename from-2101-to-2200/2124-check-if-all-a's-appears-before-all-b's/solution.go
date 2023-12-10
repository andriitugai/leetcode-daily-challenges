func checkString(s string) bool {
    firstB := strings.Index(s, "b")
    if firstB == -1 {
        firstB = len(s) + 1
    }
    return strings.LastIndex(s, "a") < firstB
}