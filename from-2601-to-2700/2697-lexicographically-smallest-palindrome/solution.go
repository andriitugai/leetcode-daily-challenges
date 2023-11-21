func makeSmallestPalindrome(s string) string {
    chars := strings.Split(s, "")
    l, r := 0, len(s) - 1
    for l < r {
        if s[l] != s[r] {
            if s[l] < s[r] {
                chars[r] = chars[l]
            } else {
                chars[l] = chars[r]
            }
        }
        l += 1
        r -= 1
    }
    return strings.Join(chars, "")
}