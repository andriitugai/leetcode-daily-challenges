func hasSpecialSubstring(s string, k int) bool {
    if k == 1 && (len(s) == 1 || s[0] != s[1]) {
        return true
    }
    prev := s[0]
    same := 1
    for i := 1; i < len(s); i ++ {
        curr := s[i]
        if curr == prev {
            same += 1
        } else {
            if same == k {
                return true
            }
            same = 1
        }
        prev = curr
    }
    if same == k {
        return true
    }
    return false
}