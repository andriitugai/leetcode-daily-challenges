func lengthOfLastWord(s string) int {
    w := strings.Split(strings.Trim(s, " "), " ")
    return len(w[len(w) - 1])
}