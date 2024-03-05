func minimumLength(s string) int {
    left, right := 0, len(s) - 1
    for left < right && s[left] == s[right] {
        for s[left] == s[left + 1] && left + 1 < right {
            left += 1
        }
        left += 1

        for s[right - 1] == s[right] && right - 1 > left {
            right -= 1
        }
        right -= 1
    }
    return right - left + 1
}